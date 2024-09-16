# Plateforme MYS3

## Initialiser le projet

### Prérequis

- Installer go
- Si utilisation VsCode installer le package Go _(Rich Go language support for Visual Studio Code)_
- Assurez-vous d'avoir Docer destop d'ouvert
- Au besoin de lancer le container api-interface tout seul, assurer vous d'avoir installer Make

  ```
  choco install make
  ```

  > _Assurez-vous de le lancer en mode administrateur_

### Intyégré les variables d'environnements au projet

Vous trouverez un `.env.exemple` à la racine du projet. Créez un fichier `.env` au même emplacement est assuréez-vous de reprendre les même termes. Ces variable d'environnements servent à la fois au lancement du projet mais aussi à sa containerisation.

```
# Need Minio settings:
S3_ENDPOINT="your-S3-endpoint"
S3_PORT= 9000
S3_ACCESSKEY= "your-S3-accesskey"
S3_SECRETKEY= "your-S3-secretkey"
S3_BUCKET= "your-minio-bucket"
DB_BOLT_PATH=my.db
```

> Si toutefois les variables ne son pas déclarées vous pouvez faire tourner le projet en local avec la command d'execution suivante, après l'installation des modules necessaires :

```powershell
go run app.go
```

Dans de telles cironcstances le programme assignera automatiquement le port `9000` ainsi que le nom et le path du store **bbolt** à la racine de `api-interface` sous le nom de `mydb`

### Cloner le repository

```
git remote add origin https://github.com/NouhaylaElfetehi/Plateforme-mys3.git
git branch -M main
git push -u origin main
```

### Installation des dépendances

```
# Clean packages
make clean-packages

# Generate go.mod & go.sum files
make requirements
```

Pour plus d'information rendez-vous sur le readme de `api-interface`

## Lancer le projet

### Lancer le projet en mode dev

Assurez vous que le package Air stable soit installer :

```
go install github.com/cosmtrek/air@v1.27.3
```

Lancer l'application en mode dev :

```
make start-dev
```

### Lancer le projet dans un container Docker

```
cd api-interface
make build
make up
marke start
```

### Lancer le projet en local

```
cd api-interface
go run app.go
```

## Utilisation des routes

### CreateBucket

| Type de requête | Route          | Params     | Body Structure            |
| --------------- | -------------- | ---------- | ------------------------- |
| PUT             | /`:buckatName` | BucketName | CreateBucketConfiguration |

#### Structure du Body pour CreateBucket

```xml
<CreateBucketConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <LocationConstraint>string</LocationConstraint>
   <Location>
      <Name>string</Name>
      <Type>string</Type>
   </Location>
   <Bucket>
      <DataRedundancy>string</DataRedundancy>
      <Type>string</Type>
   </Bucket>
</CreateBucketConfiguration>
```

##### CreateBucketConfiguration (_Obligatoire_)

Balise de niveau racine pour les paramètres CreateBucketConfiguration.

> Obligatoire : Oui

##### **[Bucket](https://docs.aws.amazon.com/fr_fr/AmazonS3/latest/API/API_CreateBucket.html#API_CreateBucket_RequestSyntax)**

Spécifie les informations sur le bucket qui sera créé.

Type : type de données [BucketInfo](https://docs.aws.amazon.com/fr_fr/AmazonS3/latest/API/API_BucketInfo.html)

Obligatoire : Non

##### Location (_non obligatoire_)

Spécifie l'emplacement où le bucket sera créé.

Pour les buckets de répertoire, le type d’emplacement est Zone de disponibilité.

Type : type de données [LocationInfo](https://docs.aws.amazon.com/fr_fr/AmazonS3/latest/API/API_LocationInfo.html)

> Obligatoire : Non

##### **[LocationConstraint](https://docs.aws.amazon.com/fr_fr/AmazonS3/latest/API/API_CreateBucket.html#API_CreateBucket_RequestSyntax)**

Spécifie la région dans laquelle le compartiment sera créé. Vous pouvez choisir une région pour optimiser la latence, minimiser les coûts ou répondre aux exigences réglementaires. Par exemple, si vous résidez en Europe, vous trouverez probablement avantageux de créer des compartiments dans la région Europe (Irlande). Pour plus d'informations, consultez [Accès à un compartiment](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro) dans le *Guide de l'utilisateur Amazon S3* .

Si vous ne spécifiez pas de région, le bucket est créé dans la région USA Est (Virginie du Nord) (us-east-1) par défaut.

Valid Values: `af-south-1 | ap-east-1 | ap-northeast-1 | ap-northeast-2 | ap-northeast-3 | ap-south-1 | ap-south-2 | ap-southeast-1 | ap-southeast-2 | ap-southeast-3 | ca-central-1 | cn-north-1 | cn-northwest-1 | EU | eu-central-1 | eu-north-1 | eu-south-1 | eu-south-2 | eu-west-1 | eu-west-2 | eu-west-3 | me-south-1 | sa-east-1 | us-east-2 | us-gov-east-1 | us-gov-west-1 | us-west-1 | us-west-2`