# Standard S3 AWS

Afin que le Bucket soit compatible avec les normes AWS S3, il est nécessaire de respecter un certain nombre de standards.

## 1. Nommage des Buckets

Le nom des buckets doit respecter des règles strictes :

Longueur du nom : Les noms de buckets doivent être composés de 3 à 63 caractères.

Caractères autorisés : Les noms de buckets ne peuvent contenir que des lettres minuscules, des chiffres, des points (.), et des tirets (-). Les lettres majuscules ne sont pas autorisées.

Début et fin du nom : Les noms de buckets doivent commencer et se terminer par une lettre ou un chiffre. Ils ne peuvent pas commencer ni finir par un point (.) ou un tiret (-).

Pas de points consécutifs : Les noms de buckets ne peuvent pas contenir deux points adjacents (par exemple, my..bucket est invalide).

Ne pas ressembler à une adresse IP : Les noms de buckets ne doivent pas être formatés comme une adresse IP (par exemple, 192.168.5.4 est invalide).

Prefixes interdits : Les noms de buckets ne doivent pas commencer par les préfixes suivants :
xn-- (préfixe réservé pour les noms de domaines internationalisés).
sthree-.
sthree-configurator-.
amzn-s3-demo- (réservé à des usages spécifiques internes AWS).

Suffixes interdits : Les noms de buckets ne doivent pas se terminer par les suffixes suivants :
-s3alias (réservé aux alias des points d'accès de bucket).
--ol-s3 (réservé aux alias des points d'accès de type Object Lambda).
.mrap (réservé pour les Multi-Region Access Points).
--x-s3 (réservé pour les directory buckets).

Unicité du nom : Les noms de buckets doivent être uniques à travers tous les comptes AWS dans toutes les Régions d’une même partition. Il existe actuellement trois partitions AWS :
aws (Régions standards).
aws-cn (Régions Chine).
aws-us-gov (AWS GovCloud).


## 2. Nommage des Objets (Fichiers)

Les objets dans S3 peuvent être nommés de manière plus flexible, mais certaines pratiques sont recommandées :

Longueur : Entre 1 et 1024 caractères.
Encodage : UTF-8.   
Caractères autorisés : Lettres minuscules et majuscules, chiffres. Tirets (-), underscores (_), points (.).
Caractères spéciaux : Les caractères spéciaux comme les espaces, accents, ou caractères spéciaux (ex : !, *, ', (, ), etc.) sont autorisés mais doivent être encodés en UTF-8 dans les URL.
Conventions recommandées :
Utiliser des noms de fichiers cohérents et descriptifs.
Éviter les caractères non imprimables ou sensibles au système de fichiers (ex : , /, ?).
Éviter de commencer les noms par un point (.), car certains systèmes peuvent les interpréter comme des fichiers cachés.


## 3. Taille des Objets (Fichiers)

AWS S3 impose certaines limitations sur la taille des fichiers :

Taille maximale pour un upload direct : 5 Go.
Taille maximale avec Multipart Upload : 5 To.
Taille minimale d’une partie dans Multipart Upload : 5 Mo (sauf la dernière partie, qui peut être plus petite).


## 4. Taille des Buckets

Nombre maximal d'objets dans un bucket : Illimité.
Pas de taille limite par défaut, mais des limites peuvent être imposées par l'utilisateur.


 ## 5. Permissions et Access Control

AWS S3 propose deux mécanismes pour gérer les permissions sur les buckets et les objets :

ACL (Access Control Lists) :

Par défaut, tous les objets sont privés.
Les ACL peuvent être définies au niveau du bucket ou de l’objet pour accorder des permissions à d’autres utilisateurs ou à des groupes prédéfinis (ex : public-read, private).
Policies (Politiques IAM) :

Les politiques IAM permettent un contrôle plus fin et basé sur des rôles pour accorder ou refuser des permissions.
Elles peuvent être définies pour un bucket entier ou pour des objets spécifiques.


## 6. Versioning

AWS S3 permet d’activer le versioning sur les buckets, ce qui permet de conserver différentes versions d'un même objet.

Chaque fois qu’un objet est modifié, une nouvelle version de cet objet est créée.
Les versions précédentes sont conservées à moins d'être explicitement supprimées.


## 7. Logging et Traçabilité

Access Logging : AWS S3 permet de conserver un journal des requêtes effectuées sur un bucket. Ces logs incluent des informations sur les requêtes effectuées, l'heure, le type de requête, etc.
Traçabilité : Intégration avec CloudTrail pour surveiller la sécurité et l'activité autour des buckets.


## 8. Classes de Stockage

AWS S3 propose plusieurs classes de stockage selon les besoins de performance et de coût :

S3 Standard : Stockage à usage général pour les données fréquemment accédées.
S3 Intelligent-Tiering : Automatisation du stockage selon les schémas d'accès.
S3 Standard-IA (Infrequent Access) : Pour les objets rarement accédés mais avec des besoins de récupération rapide.
S3 Glacier : Pour les archives long terme avec des temps de récupération plus longs.
S3 Glacier Deep Archive : Solution de stockage à moindre coût pour les données à long terme.


## 9. Consistance des Données

Consistance immédiate pour les nouveaux objets : Les objets ajoutés sont immédiatement accessibles.
Consistance éventuelle pour les suppressions et modifications : Les objets supprimés ou modifiés peuvent encore être accessibles pendant un certain temps jusqu'à ce que la modification soit propagée.


## 10. Gestion des Erreurs et Codes de Retour

Les erreurs et codes de retour doivent être conformes aux standards HTTP d'AWS S3 :

200 OK : Requête réussie.
204 No Content : Suppression réussie.
400 Bad Request : Requête mal formée (ex : taille de fichier trop grande, nom de bucket invalide).
403 Forbidden : Accès refusé.
404 Not Found : Objet ou bucket non trouvé.
500 Internal Server Error : Erreur interne du serveur.


## 11. Multipart Upload

Pour les gros fichiers, AWS S3 propose le multipart upload :

Taille minimale d’une partie : 5 Mo.
Nombre maximal de parties : 10 000.
Assemblage : Les parties téléchargées séparément sont ensuite assemblées pour former un fichier complet.


## 12. CORS (Cross-Origin Resource Sharing)

CORS permet aux applications web d’accéder aux objets S3 depuis un domaine différent. Des règles spécifiques doivent être définies pour autoriser ces requêtes inter-origines.


## 13. Gestion du Chiffrement

Chiffrement des objets au repos :
SSE-S3 : Chiffrement côté serveur avec des clés gérées par AWS.
SSE-KMS : Chiffrement avec AWS KMS (Key Management Service) pour plus de contrôle sur les clés.
SSE-C : Chiffrement avec des clés fournies par le client.
Chiffrement des objets en transit :
Utilisation du protocole HTTPS pour protéger les données lors des transferts vers et depuis S3.


## 14. Limitation du Nombre d'Objets dans un Bucket

Bien qu'il n’y ait pas de limite théorique sur le nombre d'objets, un middleware peut être ajouté pour imposer une limite personnalisée (ex : 1 million d’objets).


## 15. Gestion du Cycle de Vie des Objets

Définition de règles de cycle de vie pour les objets dans un bucket :
Transition automatique des objets vers une classe de stockage plus économique (ex : Glacier) après une certaine période.
Suppression automatique des objets après une certaine durée (ex : après 365 jours).


## 16. Vérification des Extensions de Fichiers

Valider les extensions des fichiers lors de l’upload pour s'assurer qu'elles sont autorisées (par exemple, permettre uniquement .jpg, .png, .txt).
Rejeter les fichiers avec des extensions non autorisées.

# Middlewares 

Selon les standards AWS précisés ci-dessus, l'équipe a mis en place une liste de middleware permettant de se rapprocher au maximum de ces derniers : 

- bucket_name_validation_middleware : Middleware de validation du format de nom de Bucket.
- encryption_middleware : Middleware permettant d'encrypter et décrypter les données.
- file_name_validation : Middleware de validation du format de nom de fichier.
- file_size_validation : Middleware de validation de taille de fichier pour les upload direct / Multipart.
- permissions_access_control_middleware : Middleware de contrôle des accès et permissions aux buckets et objets selon l'utilisateur.

## 1. Bucket name validation 

Ce middleware utilise des regex pour vérifier et valider le format du nom de Bucket : 
- Taille min - max 
- Prefixe, suffixe
- Double dots (..)
- etc ... 
Selon les standards AWS S3.

Exemple d'utilisation (fichier _test) : 

```Go
bucketName = "test-bucket-s3"
validator := NewBucketNameValidator()
result := validator.Validate(bucketName)
if len(result) == 1 && result[0] == "Nom de bucket valide." {
    t.Logf("Is '%s' valid? %v", test.name, result)
} else {
    t.Logf("For bucket name '%s', errors: %v", test.name, result)
}
```

se rendre dans le sous-dossier du middlware : 
```bash
go test -v ./
```

## 2. Encryption

Ce middleware permet d'encrypter et de decrypter les données.

Exemple d'utilisation (fichier _test) :
```Go
// Exemple de texte à chiffrer et déchiffrer
originalText := "This is some sensitive information"

// Créer un ResponseRecorder pour capturer les erreurs
recorder := newTestResponseWriter()

// Chiffrement
ciphertext, success := encrypt(recorder, originalText)
if !success {
    t.Fatalf("Encryption failed with status: %d", recorder.Code)
}
if len(ciphertext) == 0 {
    t.Fatal("Encryption failed: resulting ciphertext is empty")
}
t.Logf("Encrypted ciphertext: %x", ciphertext)

// Déchiffrement
recorder = newTestResponseWriter() // Reset recorder for the next operation
decryptedText, success := decrypt(recorder, ciphertext)
if !success {
    t.Fatalf("Decryption failed with status: %d", recorder.Code)
}
if decryptedText != originalText {
    t.Errorf("Decrypted text does not match original text. Got: %s, Want: %s", decryptedText, originalText)
} else {
    t.Logf("Decrypted text matches original text: %s", decryptedText)
}
```

se rendre dans le sous-dossier du middlwares : 
```bash
go test -v ./
```

## 3. File name validation 

Ce middleware utilise des regex pour vérifier et valider le format du nom de fichier : 
- taille min - max 
- Prefixe
- Suffixe
- Double dots (..)
- Patterns interdits 
- etc ...
selon les standards AWS S3.

Exemple d'utilisation (fichier _test) : 

```Go
str = "Test-file-name"

validator := NewFileNameValidator()

result := validator.Validate(str)

if len(result) == 1 && result[0] == "Nom valide." {
    t.Logf("Is '%s' valid? %v", test.name, result)
} else {
    t.Logf("For file name '%s', errors: %v", test.name, result)
}
```

se rendre dans le sous-dossier du middlwares : 
```bash
go test -v ./
```

## 4. File size validation 

Ce middleware vérifie que le fichier respecte les limties max - min pour les upload direct et multiparse : 

// Taille maximale des fichiers upload direct -> 5 Go.
// Taille maximale avec Multipart Upload : 5 To.
// Taille minimale pour les fichiers en Multipart Upload ( sauf dernier qui peut-être plus petit)

Exemple d'utilisation (fichier _test) : 

```Go
// Utilisation du middleware ValidateDirectUpload avec un handler qui renvoie 200 OK.
handler := ValidateDirectUpload(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}))

// Exécution de la requête
handler.ServeHTTP(rr, req)
```

se rendre dans le sous-dossier du middlwares : 
```bash
go test -v ./
```

## 5. Permission acces control 

Ce middleware permet de gérer les accès aux buket et objet selon ACL (Access control list).

Il se repose sur 2 entités principale : 
- User 
- Bucket 
- Bucket_object (v2)

Exemple ACL (list) : 
"bucket1": {
    "user1": "public-read",
    "user2": "private",
},
"bucket2": {
    "user3": "public-read",
    "user4": "private",
},

Exemple d'utilisation (fichier _test) : 

```Go 
// Exemple de données de test 
{
    name:               "User not in ACL",
    path:               "/bucket1/file",
    user:               "user5",
    authHeader:         "",
    expectedStatusCode: http.StatusForbidden,
}

req := httptest.NewRequest("GET", tt.path, nil)
req.Header.Set("X-User", tt.user)
req.Header.Set("Authorization", tt.authHeader)

rr := httptest.NewRecorder()
middleware := PermissionMiddleware(handler)
middleware.ServeHTTP(rr, req)

if status := rr.Code; status != tt.expectedStatusCode {
    t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatusCode)
}
```

se rendre dans le sous-dossier du middlwares : 
```bash
go test -v ./
```
