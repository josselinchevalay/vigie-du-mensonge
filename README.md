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

### Prérequis

npm installé

### Lancement

depuis le dossier frontend/vigie-du-mensonge :

```base
npm run dev
```

## Et voilà ! 
Vous pouvez vous connecter avec les identifiants suivants :

email: user@test.com -------- mdp: Test123!

Note: il est normal que vous ne puissiez ni modifier le mot de passe ni créer un compte sur la version locale.
En revanche, vous le pouvez faire sur la version live.

## Version live
Si vous voulez testez la version live -> https://vigiedumensonge.gocorp.fr