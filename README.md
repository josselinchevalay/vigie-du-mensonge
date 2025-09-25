# Vigie du mensonge

Bonjour ! J'ajoute ce petit readme pour indiquer comment run le projet rapidement, je détaillerai plus tard l'archi et les patterns de dev. 
En attendant pour les curieux, vous pouvez déjà trouver quelques infos sous .junie/guidelines.md, ainsi que sur la doc openapi et database/schema.md (à consulter sur github pour le render mermaid). 

Merci à josselin.chevalay pour ta contribution sur docker <3

## BDD et Backend

### Prérequis

docker installé

### Lancement

depuis la racine du projet :

```bash
docker compose up -d --build
```

### Shutdown

```bash
docker compose down --rmi all
```

## Frontend

depuis le dossier frontend/vigie-du-mensonge :

```base
npm run dev
```

## Et voilà ! Si vous voulez testez la version live -> https://vigiedumensonge.gocorp.fr