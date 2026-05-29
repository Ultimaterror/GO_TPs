# Cours Jour 2

Ce qui font du Go : des Gofers  
Pas de conversion implicite  
La valeur initiale d’une variable est « valeur zéro »

- int : 0
- float : 0.0
- string : ""
- bool : false
- slice : nil
- function : nil

Slice : Tableau avec taille modifiable  
Façons de déclarer une variable

- Type défini
  - var name string
- Inférence de type
  - var ville = ‘Paris’
- Courte (Seulement dans les fonctions)
  - nom := ‘Bob’
  - 2 variables d’un coup
    - x, y := 10, 20
- Multiples
  - ```
    var (
        serveur string = « localhost »
        port int = 8000
    )
    ```

Echange de valeur (pas besoin de variable intermédiaire)

- x, y = y, x

Ne compile pas si une variable n’est pas utilisée  
iota commence à 0  
Une fonction ne pas être surchargé  
Retour multiples important en GO

- Classique
  - Resultat & erreur
    - Permet de toujours vérifier l’erreur

Ignorer une valeur

- Utiliser underscore \_
  - Exemple : res, \_

Retours nommés  
En GO, l’erreur est une valeur à gérer  
Closure : Une fonction qui se souvient de l’environnement dans lequel elle a été créée  
Une seule boucle : for  
SwitchCase « fallthrough » : Permet de passer à une étape suivante même si une condition est remplie  
Le « fallthrough » existe ou pas en fonction des langages et quand il existe c’est souvent implice (pas indiqué)  
slice : sac qui peut être agrandi (taille dynamique)

- Longueur
- Capacité (longueur max)

map : dictionnaire cle/valeur  
struct : objet en JS

- 3 façons d’initialiser
  - Positionnelle
    - Dans l’ordre des clés
    - Fragile
  - Nommée
    - À utiliser
  - Affectation
    - Ex : Obj.Nom

Composition : Assembler des pièces pour composer un objet final  
tags : chaines de caractères entre backticks « ` »  
la casse est importante pour les struct

- Nom ou func Validate
  - Accessible
- nom ou func validate
  - Non exporté

defer : Executer avant le return même s’il y a une erreur après

- Ordre « Last In First Out » (LIFO)

Packages standard incontournables

- fmt
  - Formatage et affichage
- strings
  - Manipulation de strings
- strconv
  - Conversions string  types
- math
  - Fonctions mathématiques
- sort
  - Tri
- time
  - Date et heure
