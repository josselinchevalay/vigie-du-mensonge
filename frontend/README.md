

# Frontend – Vigie du mensonge

Ce dossier contient le code du frontend du projet.  
C’est une application **React + TypeScript** construite avec **Vite**, qui consomme l’API backend et fournit l’interface utilisateur.

---

## 🚀 Démarrage rapide

[Voir le README.md à la racine du projet](../README.md#lancement-de-lapplication)

Note: un bindmount est setup pour le **hot reload** du frontend dans le docker-compose.

## 🛠️ Stack technique

- **React + TypeScript** – Framework UI
- **Vite** – Outil de build
- **TanStack Router** – Routing type-safe
- **TanStack Query** – Gestion du server state (fetch + cache + revalidation) avec invalidation automatique, mutations, pagination et retries.
- **TanStack Store** – Gestion d’état légère
- **Ky** – Client HTTP
- **Tailwind CSS** – Styles utilitaires
- **Shadcn UI** – Composants UI pré-construits
- **Vitest + Testing Library** – Tests

### Point d'attention - Tanstack Router

L’application utilise le **file-based routing** fourni par TanStack Router.

👉 Fonctionnement :
- Chaque fichier créé dans `src/routes/` correspond automatiquement à une route.
- Le nom du fichier définit l’URL de la route.
- Lors du build, TanStack Router génère automatiquement l’arborescence des routes.

### Exemple : ajouter une route `/test`

1. Arrêter l’application si elle tourne encore.
2. Aller dans le dossier `src/routes/`.
3. Créer un fichier `test.tsx` (vide).
4. Relancer un build avec :
   ```sh
   npm run build
   ```
5. Relancez l'app, le code a été généré et la route est maintenant accessible sur [http://localhost:5173/test](http://localhost:5173/test).

⚠️ Inutile de modifier manuellement la configuration du routeur : TanStack Router s’occupe de générer les routes automatiquement à partir des fichiers présents.
 

---

## 📂 Organisation du code

```
src/
├── routes/           # Pages et routes TanStack Router
├── core/
│   ├── shadcn/       # Composants générés via Shadcn
│   ├── components/   # Composants UI personnalisés
│   ├── dependencies/ # Services partagés (API, config…)
│   └── models/       # Modèles TypeScript
└── index.css         # Styles globaux Tailwind
```

Autres dossiers :
- `public/` : assets statiques
- `ops/` : configuration nginx (docker)
- `dist/` : build final

---

## ✅ Tests

Les tests sont écrits avec **Vitest** et **@testing-library/react**.  

Lancer tous les tests :
```sh
npm test
```

---

## 📌 Bonnes pratiques de contribution

- Ajouter les nouvelles pages dans `src/routes/`
- Les composants réutilisables vont dans `src/core/components/`
- Les modèles TypeScript dans `src/core/models/`
- Les services (ex. appels API via Ky) dans `src/core/dependencies/`
- Les composants générés via Shadcn restent isolés dans `src/core/shadcn/`
- Toujours typer les props et données avec **TypeScript**

---

## 🤝 Contribution

Toutes les suggestions sont bienvenues.  
Le projet reste ouvert aux améliorations proposées par la communauté.