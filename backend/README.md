

# Backend – Vigie du mensonge

Ce dossier contient l’API du projet.  
Elle expose des routes HTTP utilisées par le frontend pour gérer les articles, les utilisateurs et la modération.

---

## 🚀 Démarrage rapide

[Voir le README.md à la racine du projet](../README.md#lancement-de-lapplication)

## 🛠️ Stack technique

- **Go** – langage compilé, rapide, typé
- **Fiber** – framework web (analogue à Express en Node.js)
- **GORM** – ORM (analogue à Sequelize)
- **PostgreSQL** – base de données relationnelle
- **OpenAPI** – documentation automatique de l’API
- **Docker** – pour déploiement et environnement isolé
- **testcontainers** – pour les tests d’intégration

---

## 📂 Organisation du code

### `main.go`
Point d’entrée de l’application : démarre le serveur Fiber et enregistre les routes.

### `/api`
Chaque fonctionnalité a son propre dossier, contenant :
- `handler.go` – endpoints HTTP (contrôleurs)
- `service.go` – logique métier
- `repository.go` – accès DB spécifique (optionnel)
- `dto.go` – structures d’entrée/sortie
- `*_test.go` – tests unitaires
- `integration_test.go` – test d’intégration obligatoire

### `/core`
Fonctionnalités transverses :
- **dto/** : objets partagés pour les réponses
- **logger/** : wrapper de log
- **jwt_utils/** et **hmac_utils/** : sécurité
- **dependencies/** : connexions DB, mailer
- **models/** : définitions GORM des tables (User, Article, Politician…)
- **env/** : variables d’environnement & config
- **locals/** : données stockées dans le contexte Fiber (user authentifié, tokens…)
- **fiberx/** : extensions Fiber
- **validation/** : règles de validation

### `/test_utils`
Utilitaires pour simplifier l’écriture de tests.

---

## 🔄 Cycle de vie d’une requête

1. Un client envoie une requête HTTP → **handler.go**
2. Le handler appelle le **service.go** correspondant
3. Le service applique la logique métier et appelle éventuellement un **repository.go**
4. Le repository interagit avec la **base de données** via **GORM**
5. Réponse renvoyée en JSON (formaté par un **DTO**)

---

## ✅ Tests

- **Unit tests** : présents dans `handler_test.go` et `service_test.go`
- **Integration tests** : un fichier `integration_test.go` par feature, utilisant Testcontainers pour démarrer une vraie DB PostgreSQL

Lancer tous les tests :
```sh
go test ./...
```

---

## 📌 Bonnes pratiques de contribution

- Ajouter un **integration_test.go** pour chaque nouvelle route
- Respecter la séparation : `handler` (I/O), `service` (logique métier), `repository` (DB)
- Toujours utiliser des **DTOs** pour les entrées/sorties
- Logger les erreurs avec `core/logger`
- Ne jamais coder de mot de passe/token en dur → utiliser `core/env`

---

## 🔮 Roadmap backend

- Migration possible de GORM vers **sqlc** pour de meilleures performances (gorm est pratique mais gourmand à cause de la reflection)
- Amélioration du système de logs
- Ajout d’outils de monitoring

---

## 🤝 Contribution

Tout est ouvert aux suggestions !  
N’hésitez pas à proposer des améliorations, l’architecture est conçue pour évoluer selon les retours de la communauté.