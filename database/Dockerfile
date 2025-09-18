# Utiliser l'image officielle PostgreSQL
FROM postgres:16

# Définir les variables d'environnement pour PostgreSQL
ENV POSTGRES_DB=vdm_db
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=root

# Copier le fichier d'initialisation SQL dans le répertoire d'initialisation de PostgreSQL
COPY schema.sql /docker-entrypoint-initdb.d/

# Exposer le port PostgreSQL par défaut
EXPOSE 5432