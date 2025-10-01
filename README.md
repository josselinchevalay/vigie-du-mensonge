# Vigie du mensonge

https://vigiedumensonge.gocorp.fr/

## IMPORTANT

Ce repo est à l'origine une initiative personnelle de ma part en tant que membre du discord de la
communauté de **Clemovitch** et n'a rien d'**officiel**.

Il s'agit simplement d'une proposition alternative qui vise à mettre en place une plateforme permettant
de référencer les mensonges, les contre-vérités, les manipulations, et la désinformation des gouvernements
français en attendant qu'une version officielle soit mise en place. Le cas échéant, je ne vois aucun obstacle
majeur au portage des articles qui auront été rédigés sur ce projet vers le projet officiel.

## Vue d'ensemble

Ce README est destiné au grand public et ne contient donc aucun détail technique
(en dehors de la section de [lancement de l'application](#lancement-de-lapplication)). Pour cela, je vous invite à vous
référer aux liens suivants :

- [FRONTEND](./frontend/README.md)
- [BACKEND](./backend/README.md)
- [DATA_IMPORT](./data_import/README.md) - module d'import de jeux de données data.gouv
- [DATABASE](./database/README.md)

## Fonctionnalités actuelles

### Visiteurs

- Inscription & connexion sécurisée (sessions persistantes)
- Consultation des articles publiés ➡️ [page d'accueil](./docs/screenshots/home.png)
- Consultation du contenu détaillé d'un article ➡️ [page d'un article](./docs/screenshots/article_visitor.png)

### Rédacteurs

- Accéder à un **espace rédacteur**
  ➡️ [sur le web](./docs/screenshots/access_redactor_web.png) / [sur mobile](./docs/screenshots/access_redactor_mobile.png)
- Consulter ses propres articles dans l'espace rédacteur
  ➡️ [page d'accueil de l'espace rédacteur](./docs/screenshots/redactor_home.png)
- Créer un nouveau brouillon d'article
  ➡️ [formulaire de création d'un article](./docs/screenshots/redactor_article_form.png)
- Modification avec versioning des articles et notes de la modération
  ➡️ [consulter/modifier un de ses articles](./docs/screenshots/redactor_articles_version.png)

### Modérateurs

- Accéder à un **espace modérateur**
  ➡️ [sur le web](./docs/screenshots/access_moderator_web.png) / [sur mobile](./docs/screenshots/access_moderator_mobile.png)
- Consulter les articles sous sa propre modération
  ➡️ [page d'accueil de l'espace modérateur](./docs/screenshots/moderator_home.png)
- Approuver/refuser la publication d'articles sous sa modération
  ➡️ [formulaire de review du modérateur](./docs/screenshots/moderator_review_form.png)
- Consulter les articles en attente de modération
  ➡️ [page des articles en attente de modération](./docs/screenshots/moderator_pending_articles.png)
- Revendiquer la modération d'un article
  ➡️ [page d'un article en attente de modération](./docs/screenshots/moderator_claims_pending_article.png)


## Fonctionnalités à venir 

- Espace administrateur pour : 
  - rechercher un utilisateur à partir de son nom d'utilisateur dans une barre de recherche
  - afficher le profil d'un utilisateur avec son historique d'activités (rédaction, modération d'articles...)
  - ajouter / retirer les autorisations d'un utilisateur (rédacteur, modérateur)

- Page dédiée au top des politiques les plus mentionnés dans les articles
- Page dédiée à un politique avec l'ensemble des articles le concernant ainsi que ses 
différentes fonctions au sein des gouvernements


## Lancement de l'application

### Prérequis

- **Docker** et **Docker Compose** installés sur votre machine
- Aucun service en cours d’exécution sur les ports utilisés par défaut
    - 5173 pour le frontend
    - 8080 pour le backend
    - 5432 pour la base de données

### Démarrage

À la racine du projet, lancer simplement :

```sh
docker compose up --build
```

### Openapi

L'application expose un openapi sur http://localhost:8080/docs pour la documentation de l'API.

## Pour conclure

Toutes les suggestions sont les bienvenues, et l’ensemble du projet reste entièrement revisitable en fonction des retours de la communauté.
Vous pouvez ouvrir une GitHub issue pour signaler un bug ou suggérer une amélioration / modification / nouvelle feature.