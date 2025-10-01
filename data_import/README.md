

# Data Import – Vigie du mensonge

Ce module permet d’initialiser et de peupler la base de données PostgreSQL avec des données nécessaires au bon fonctionnement de l’application (**politicians, governments, occupations**).

Les fichiers CSV contenant les données proviennent du site [data.gouv.fr](https://www.data.gouv.fr/datasets/historique-des-gouvernements-de-la-veme-republique/).

---

## 🚀 Démarrage rapide

Ce module est executé automatiquement via le docker-compose à la racine du projet.

## ⚙️ Fonctionnement

- Se connecte à la base PostgreSQL via les [variables d’environnement](../.db.env).
- Insère les données issues de [governments.csv](governments.csv), [occupations.csv](occupations.csv), [presidents.csv](presidents.csv) dans les tables suivantes:
  - politicians
  - governments
  - occupations

---

## 🤝 Contribution

Ce module reste simple et peut évoluer (par ex. importer de nouvelles données de référence).  
N’hésitez pas à proposer des améliorations pour enrichir le jeu de données initial.