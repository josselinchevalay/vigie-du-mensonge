package controller

import (
	"context"
	"fmt"

	appv1 "github.com/geo1796/vigie-du-mensonge/k8s/operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// VigieAppReconciler reconciles a VigieApp object
type VigieAppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=vigie.vigie.local,resources=vigieapps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=vigie.vigie.local,resources=vigieapps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete

func (r *VigieAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var app appv1.VigieApp
	if err := r.Get(ctx, req.NamespacedName, &app); err != nil {
		if errors.IsNotFound(err) {
			// Ressource supprimée
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// 1️⃣ Crée/Met à jour le Deployment Backend
	backendDeployment := &appsv1.Deployment{}
	backendDeploymentName := fmt.Sprintf("%s-backend", app.Name)
	err := r.Get(ctx, types.NamespacedName{Name: backendDeploymentName, Namespace: req.Namespace}, backendDeployment)
	if err != nil && errors.IsNotFound(err) {
		backendDeployment = r.buildDeployment(app, "backend", app.Spec.Backend.Image, app.Spec.Backend.Replicas)
		if err := r.Create(ctx, backendDeployment); err != nil {
			return ctrl.Result{}, err
		}
		logger.Info("Backend Deployment created", "name", backendDeploymentName)
	}

	// 2️⃣ Crée/Met à jour le Deployment Frontend
	frontendDeployment := &appsv1.Deployment{}
	frontendDeploymentName := fmt.Sprintf("%s-frontend", app.Name)
	err = r.Get(ctx, types.NamespacedName{Name: frontendDeploymentName, Namespace: req.Namespace}, frontendDeployment)
	if err != nil && errors.IsNotFound(err) {
		frontendDeployment = r.buildDeployment(app, "frontend", app.Spec.Frontend.Image, app.Spec.Frontend.Replicas)
		if err := r.Create(ctx, frontendDeployment); err != nil {
			return ctrl.Result{}, err
		}
		logger.Info("Frontend Deployment created", "name", frontendDeploymentName)
	}

	// 3️⃣ Crée le Service pour Backend et Frontend
	// (exemple simple, on ne gère que le backend ici)
	service := &corev1.Service{}
	serviceName := fmt.Sprintf("%s-backend", app.Name)
	err = r.Get(ctx, types.NamespacedName{Name: serviceName, Namespace: req.Namespace}, service)
	if err != nil && errors.IsNotFound(err) {
		service = r.buildService(app, "backend", 8080)
		if err := r.Create(ctx, service); err != nil {
			return ctrl.Result{}, err
		}
		logger.Info("Backend Service created", "name", serviceName)
	}

	// 4️⃣ Crée un Ingress si activé
	if app.Spec.Ingress.Enabled {
		ing := &networkingv1.Ingress{}
		ingName := fmt.Sprintf("%s-ingress", app.Name)
		err = r.Get(ctx, types.NamespacedName{Name: ingName, Namespace: req.Namespace}, ing)
		if err != nil && errors.IsNotFound(err) {
			ing = r.buildIngress(app, 80, app.Spec.Ingress.Domain)
			if err := r.Create(ctx, ing); err != nil {
				return ctrl.Result{}, err
			}
			logger.Info("Ingress created", "name", ingName)
		}
	}

	// --- DATABASE PVC ---
	dbPVCName := fmt.Sprintf("%s-database-pvc", app.Name)
	pvc := &corev1.PersistentVolumeClaim{}
	err = r.Get(ctx, types.NamespacedName{Name: dbPVCName, Namespace: app.Namespace}, pvc)
	if err != nil && errors.IsNotFound(err) {
		pvc = &corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:      dbPVCName,
				Namespace: app.Namespace,
			},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes: []corev1.PersistentVolumeAccessMode{
					corev1.ReadWriteOnce,
				},
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse(app.Spec.Database.Storage),
					},
				},
			},
		}
		if err := r.Create(ctx, pvc); err != nil {
			return ctrl.Result{}, err
		}
		logger.Info("Database PVC created", "name", dbPVCName)
	}

	// --- DATABASE DEPLOYMENT ---
	dbDeploymentName := fmt.Sprintf("%s-database", app.Name)
	dbDeployment := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: dbDeploymentName, Namespace: app.Namespace}, dbDeployment)
	if err != nil && errors.IsNotFound(err) {
		dbDeployment = &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      dbDeploymentName,
				Namespace: app.Namespace,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: int32Ptr(1),
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"app": dbDeploymentName},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{"app": dbDeploymentName},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  "database",
								Image: "postgres:15",
								Ports: []corev1.ContainerPort{
									{ContainerPort: 5432},
								},
								Env: []corev1.EnvVar{
									{
										Name:  "POSTGRES_USER",
										Value: app.Spec.Database.User,
									},
									{
										Name: "POSTGRES_PASSWORD",
										ValueFrom: &corev1.EnvVarSource{
											SecretKeyRef: &corev1.SecretKeySelector{
												LocalObjectReference: corev1.LocalObjectReference{
													Name: app.Spec.Database.PasswordSecretRef,
												},
												Key: "password",
											},
										},
									},
								},
								VolumeMounts: []corev1.VolumeMount{
									{
										Name:      "db-storage",
										MountPath: "/var/lib/postgresql/data",
									},
								},
							},
						},
						Volumes: []corev1.Volume{
							{
								Name: "db-storage",
								VolumeSource: corev1.VolumeSource{
									PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
										ClaimName: dbPVCName,
									},
								},
							},
						},
					},
				},
			},
		}
		if err := r.Create(ctx, dbDeployment); err != nil {
			return ctrl.Result{}, err
		}
		logger.Info("Database Deployment created", "name", dbDeploymentName)
	}


	// --- DATA IMPORT ---
	if app.Spec.DataImport.Enabled {
		dataDeploymentName := fmt.Sprintf("%s-dataimport", app.Name)
		dataDeployment := &appsv1.Deployment{}
		err = r.Get(ctx, types.NamespacedName{Name: dataDeploymentName, Namespace: app.Namespace}, dataDeployment)
		if err != nil && errors.IsNotFound(err) {
			dataDeployment = r.buildDeployment(app, "dataimport", app.Spec.DataImport.Image, 1)
			if err := r.Create(ctx, dataDeployment); err != nil {
				return ctrl.Result{}, err
			}
			logger.Info("DataImport Deployment created", "name", dataDeploymentName)
		}
	}

	// --- FRONTEND & BACKEND ---
	r.ensureComponent(ctx, &app, "frontend", app.Spec.Frontend.Image, app.Spec.Frontend.Replicas)
	r.ensureComponent(ctx, &app, "backend", app.Spec.Backend.Image, app.Spec.Backend.Replicas)

	// --- SERVICE ---
	r.ensureService(ctx, &app, "backend", 8080)
	r.ensureService(ctx, &app, "frontend", 80)

	// --- INGRESS ---
	if app.Spec.Ingress.Enabled {
		r.ensureIngress(ctx, &app)
	}

	// 5️⃣ Met à jour le Status
	app.Status.Ready = true
	app.Status.BackendPort = 8080
	app.Status.FrontendPort = 80
	app.Status.Message = "VigieApp deployed successfully"
	if err := r.Status().Update(ctx, &app); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// --- Helpers ---
func (r *VigieAppReconciler) ensureComponent(ctx context.Context, app *appv1.VigieApp, name, image string, replicas int32) {
	deploy := &appsv1.Deployment{}
	deployName := fmt.Sprintf("%s-%s", app.Name, name)
	err := r.Get(ctx, types.NamespacedName{Name: deployName, Namespace: app.Namespace}, deploy)
	if err != nil && errors.IsNotFound(err) {
		deploy = r.buildDeployment(*app, name, image, replicas)
		_ = r.Create(ctx, deploy)
	}
}

func (r *VigieAppReconciler) ensureService(ctx context.Context, app *appv1.VigieApp, name string, port int32) {
	svc := &corev1.Service{}
	svcName := fmt.Sprintf("%s-%s", app.Name, name)
	err := r.Get(ctx, types.NamespacedName{Name: svcName, Namespace: app.Namespace}, svc)
	if err != nil && errors.IsNotFound(err) {
		svc = r.buildService(*app, name, port)
		_ = r.Create(ctx, svc)
	}
}

func (r *VigieAppReconciler) ensureIngress(ctx context.Context, app *appv1.VigieApp) {
	ing := &networkingv1.Ingress{}
	ingName := fmt.Sprintf("%s-ingress", app.Name)
	err := r.Get(ctx, types.NamespacedName{Name: ingName, Namespace: app.Namespace}, ing)
	if err != nil && errors.IsNotFound(err) {
		ing = r.buildIngress(*app, 80, app.Spec.Ingress.Domain)
		_ = r.Create(ctx, ing)
	}
}

func int32Ptr(i int32) *int32 { return &i }

func (r *VigieAppReconciler) buildDeployment(app appv1.VigieApp, name, image string, replicas int32) *appsv1.Deployment {
	labels := map[string]string{"app": name, "vigieapp": app.Name}
	return &appsv1.Deployment{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      fmt.Sprintf("%s-%s", app.Name, name),
			Namespace: app.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{MatchLabels: labels},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: ctrl.ObjectMeta{Labels: labels},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  name,
							Image: image,
							Ports: []corev1.ContainerPort{{ContainerPort: 8080}},
						},
					},
				},
			},
		},
	}
}

func (r *VigieAppReconciler) buildService(app appv1.VigieApp, name string, port int32) *corev1.Service {
	labels := map[string]string{"app": name, "vigieapp": app.Name}
	return &corev1.Service{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      fmt.Sprintf("%s-%s", app.Name, name),
			Namespace: app.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Port:       port,
					TargetPort: intstr.FromInt(int(port)),
				},
			},
		},
	}
}

func (r *VigieAppReconciler) buildIngress(app appv1.VigieApp, port int32, host string) *networkingv1.Ingress {
	pathType := networkingv1.PathTypePrefix
	return &networkingv1.Ingress{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      fmt.Sprintf("%s-ingress", app.Name),
			Namespace: app.Namespace,
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: host,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathType,
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: fmt.Sprintf("%s-backend", app.Name),
											Port: networkingv1.ServiceBackendPort{Number: port},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *VigieAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.VigieApp{}).
		Complete(r)
}
