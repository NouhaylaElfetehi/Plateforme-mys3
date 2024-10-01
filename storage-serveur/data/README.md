# Serveur de Stockage

Ce projet implémente un serveur de stockage simple en Go, capable de gérer le téléchargement et la récupération de fichiers. Les fichiers sont stockés dans un répertoire spécifique et peuvent être recherchés efficacement à l'aide d'une recherche binaire.

## Fonctionnalités

- **Téléchargement de fichiers** : Les fichiers peuvent être téléchargés via une requête HTTP POST.
- **Récupération de fichiers** : Les fichiers peuvent être récupérés via une requête HTTP GET.
- **Recherche binaire** : Utilisation d'un algorithme de recherche binaire pour rechercher efficacement les fichiers par leur nom.

## Choix Techniques

### 1. Utilisation de la Recherche Binaire

Nous avons choisi d'utiliser une recherche binaire pour rechercher les fichiers par leur nom. La recherche binaire est un algorithme efficace pour trouver un élément dans une liste triée, avec une complexité temporelle de O(log n). Cela permet de rechercher rapidement un fichier spécifique parmi de nombreux fichiers stockés.

### 2. Chargement Initial des Fichiers

Au démarrage du serveur, nous chargeons et trions la liste des fichiers présents dans le répertoire `./data/`. Cela permet de maintenir une liste triée des fichiers en mémoire, ce qui est nécessaire pour la recherche binaire.

### 3. Maintien de la Liste Triée

Lors de l'ajout d'un nouveau fichier, nous mettons à jour la liste des fichiers et la trions à nouveau. Cela garantit que la liste des fichiers reste triée en permanence, ce qui est essentiel pour la recherche binaire.

### 4. Verrouillage Concurent

Nous utilisons `sync.RWMutex` pour gérer les accès concurrents à la liste des fichiers. Les verrous en lecture sont utilisés pour les opérations de recherche, tandis que les verrous en écriture sont utilisés pour les opérations de mise à jour de la liste des fichiers. Cela permet de minimiser les sections critiques et d'améliorer la concurrence.

## Exemple de Code

### Recherche Binaire

````go
func binarySearch(name string) string {
    low, high := 0, len(fileNames)-1
    for low <= high {
        mid := (low + high) / 2
        if fileNames[mid] == name {
            return fileNames[mid]
        } else if fileNames[mid] < name {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return ""
}

## Exécution du Serveur

- Pour exécuter le serveur, utilisez la commande suivante :

```shell
go run main.go
````

## Docker

Un Dockerfile est fourni pour exécuter le serveur dans un conteneur Docker. Pour construire et exécuter l'image Docker, utilisez les commandes suivantes :

```shell
docker build -t storage-server .
docker run -d -p 8081:8081 -v /path/to/local/data:/root/data storage-server
```

Remplacez /path/to/local/data par le chemin local où vous souhaitez stocker les fichiers.

## Détail des fonctionnalités

### Fonction `init`

**Rôle** : Initialiser la liste des noms de fichiers au démarrage du serveur.

**Explication** :

- Cette fonction est automatiquement appelée au démarrage du programme.
- Elle appelle la fonction `loadFileNames` pour charger et trier la liste des fichiers présents dans le répertoire `./data/`.

### Fonction `loadFileNames`

**Rôle** : Charger et trier la liste des noms de fichiers présents dans le répertoire `./data/`.

**Explication** :

- Utilise `filepath.Glob` pour lister tous les fichiers dans le répertoire `./data/`.
- Extrait le nom de base de chaque fichier et les ajoute à la liste `fileNames`.
- Trie la liste `fileNames` pour permettre une recherche binaire efficace.

### Fonction `binarySearch`

**Rôle** : Rechercher un fichier par son nom dans la liste triée `fileNames` en utilisant l'algorithme de recherche binaire.

**Explication** :

- Initialise deux indices, `low` et `high`, pour représenter les limites de la recherche.
- Utilise une boucle pour diviser la liste en deux parties à chaque étape et comparer l'élément recherché avec l'élément au milieu de la liste.
- Si l'élément est trouvé, il est retourné. Sinon, la recherche continue dans la moitié appropriée de la liste.
- Retourne une chaîne vide si l'élément n'est pas trouvé.

### Fonction `uploadFile`

**Rôle** : Gérer le téléchargement de fichiers via une requête HTTP POST et mettre à jour la liste des noms de fichiers.

**Explication** :

- Récupère le fichier et ses métadonnées à partir de la requête.
- Crée un nouveau fichier dans le répertoire `./data/` avec le nom du fichier téléchargé.
- Copie le contenu du fichier téléchargé dans le nouveau fichier.
- Ajoute le nom du fichier à la liste `fileNames` et trie la liste pour maintenir l'ordre nécessaire à la recherche binaire.
- Utilise un verrou (`mutex`) pour garantir que les opérations de mise à jour de la liste des fichiers sont thread-safe.
- Répond avec un statut HTTP 201 pour indiquer que le fichier a été créé avec succès.

### Fonction `downloadFile`

**Rôle** : Gérer la récupération de fichiers via une requête HTTP GET en utilisant la recherche binaire pour trouver le fichier par son nom.

**Explication** :

- Récupère le paramètre `name` de la requête pour obtenir le nom du fichier à télécharger.
- Utilise la fonction `binarySearch` pour rechercher le fichier dans la liste triée `fileNames`.
- Si le fichier est trouvé, il est ouvert et son contenu est servi au client avec le type MIME approprié.
- Utilise un verrou en lecture (`mutex.RLock`) pour garantir que les opérations de recherche sont thread-safe.
- Répond avec un statut HTTP 404 si le fichier n'est pas trouvé.

### Fonction `main`

**Rôle** : Configurer les routes HTTP et démarrer le serveur.

**Explication** :

- Configure les routes HTTP pour les endpoints `/upload` et `/download`.
- Démarre le serveur HTTP sur le port 8081.
- Utilise `log.Fatal` pour enregistrer toute erreur qui pourrait survenir lors du démarrage du serveur.

### Critique sur l'approche actuelle

#### Avantages de l'approche actuelle

- **Recherche rapide** : La recherche binaire a une complexité temporelle de \(O(\log n)\), ce qui est très efficace pour des recherches dans une liste triée.
- **Simplicité** : La solution est simple à implémenter et à comprendre, ce qui facilite la maintenance.

#### Limites de l'approche actuelle

- **Scalabilité** : La liste des fichiers est maintenue en mémoire, ce qui peut devenir problématique avec un très grand nombre de fichiers.
- **Synchronisation** : La gestion des verrous pour les accès concurrents peut devenir un goulot d'étranglement avec un grand nombre de requêtes simultanées.
- **Distribution** : La synchronisation des données entre plusieurs réplicas peut devenir complexe et coûteuse en termes de performance.

#### Optimisations possibles pour des volumes de données importants

- **Base de données** : Utiliser une base de données (SQL ou NoSQL) pour stocker les métadonnées des fichiers. Cela permet de gérer des volumes de données beaucoup plus importants et de bénéficier des optimisations de recherche et de gestion de transactions offertes par les bases de données.
- **Indexation** : Utiliser des index pour accélérer les recherches dans la base de données.
- **Cache** : Mettre en place un système de cache (comme Redis) pour stocker les résultats des recherches fréquentes et réduire la charge sur la base de données.
- **Partitionnement** : Diviser les données en partitions pour répartir la charge entre plusieurs serveurs et améliorer la scalabilité.
- **Consistent Hashing** : Utiliser un algorithme de hachage consistant pour répartir les fichiers entre plusieurs nœuds de manière équilibrée.

#### Architecture distribuée avec réplicas

Pour une architecture avec trois réplicas (deux pour les opérations de lecture/écriture et un pour la sauvegarde), voici quelques points à considérer :

- **Consistance** : Assurer la consistance des données entre les réplicas. Utiliser des protocoles comme Raft ou Paxos pour la réplication et la gestion des consensus.
- **Disponibilité** : Assurer une haute disponibilité en répartissant les requêtes entre les réplicas et en utilisant des mécanismes de basculement en cas de défaillance d'un nœud.
- **Partitionnement** : Répartir les données entre les réplicas pour éviter les goulots d'étranglement et améliorer la scalabilité.
- **Synchronisation** : Mettre en place des mécanismes de synchronisation efficaces pour maintenir les réplicas à jour, comme la réplication asynchrone ou les journaux de transactions.

#### Potentiel d'efficacité

L'efficacité de cette architecture dépendra de plusieurs facteurs, notamment :

- **Volume de données** : Pour des volumes de données très importants (plusieurs téraoctets), une base de données distribuée et partitionnée sera plus efficace qu'une liste en mémoire.
- **Nombre de requêtes** : Pour un grand nombre de requêtes simultanées, l'utilisation de caches et de mécanismes de partitionnement améliorera les performances.
- **Latence** : La latence des opérations de synchronisation entre les réplicas peut affecter les performances globales. Utiliser des techniques de réplication efficaces est crucial.

## Priorités de correction

### Mise en place d'un système de cache

- **Action** : Utiliser un système de cache (comme Redis) pour stocker les résultats des recherches fréquentes.
- **Bénéfice** : Réduit la charge sur la base de données et améliore les performances des requêtes fréquentes.

### Optimisation des verrous

- **Action** : Réduire la granularité des verrous ou utiliser des mécanismes de synchronisation plus efficaces.
- **Bénéfice** : Améliore les performances en cas de forte concurrence et réduit les risques de blocages.

### Partitionnement des données

- **Action** : Diviser les données en partitions pour répartir la charge entre plusieurs serveurs.
- **Bénéfice** : Améliore la scalabilité et les performances en répartissant la charge de manière équilibrée.

### Consistent Hashing pour la distribution des fichiers

- **Action** : Utiliser un algorithme de hachage consistant pour répartir les fichiers entre plusieurs nœuds.
- **Bénéfice** : Assure une répartition équilibrée des fichiers et améliore la scalabilité du système distribué.

### Protocoles de réplication et de consensus

- **Action** : Mettre en place des protocoles comme Raft ou Paxos pour la réplication et la gestion des consensus.
- **Bénéfice** : Assure la consistance des données entre les réplicas et améliore la fiabilité du système.
