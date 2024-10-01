# Étape de construction
FROM golang:1.18 AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers du projet
COPY . .

# Télécharger les dépendances et construire l'application
RUN go mod tidy
RUN go build -o storage-server .

# Étape finale
FROM alpine:latest

# Installer les certificats CA
RUN apk --no-cache add ca-certificates

# Définir le répertoire de travail
WORKDIR /root/

# Copier le binaire construit depuis l'étape de construction
COPY --from=builder /app/storage-server .

# Créer le répertoire de données
RUN mkdir -p /root/data

# Exposer le port
EXPOSE 8081

# Définir le volume pour le répertoire de données
VOLUME ["/root/data"]

# Commande d'exécution
CMD ["./storage-server"]