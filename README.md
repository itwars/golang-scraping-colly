Différents exemples de web scraping en utilisant le framework Colly (Golang).
- Le premier exemple *freelance-info* récupère les tarifs journaliers des freelances :
```json
{"Administrateur BD":540,"Administrateur ERP":510,"Administrateur produits":470,"Administrateur réseaux":390,"Administrateur système":430,"Analyste":450,"Analyste d'exploitation":340,"Analyste programmeur":390,"Analyste réalisateur":410,"Architecte":620,"Architecte réseaux":590,"Assistant à maîtrise d'ouvrage":570,"Auditeur":660,"Chef de projet":570,"Concepteur BD":430,"Concepteur multimédia":380,"Concepteur télématique":450,"Consultant":600,"Consultant fonctionnel":650,"Consultant technique":540,"Consultant technique et formateur":640,"Directeur de projet":760,"Directeur informatique":790,"Développeur":410,"Expert":650,"Formateur":510,"Infographiste":330,"Ingénieur d'exploitation":430,"Ingénieur d'études":440,"Ingénieur de production":470,"Ingénieur réseaux":480,"Ingénieur système":500,"Maquettiste PAO":320,"Pupitreur/Pilote":300,"Responsable d'exploitation":550,"Responsable maintenance":390,"Rédacteur technique":340,"Support utilisateurs":310,"Technicien d'exploitation":280,"Technicien micro / réseaux":270,"Webmaster":320}
```
- Le second exemple *meteo* fournit des prévisions et informations de différents sites météos et surf :
```
Département Girondes
Date du jour          :  Lundi 27 août
Le soleil se lève à   :  07h18
Le soleil se couche à :  20h49
Saint(e) du jour      :  Sainte Monique
Tendance du temps     :  Éclaircies
Température           :  15°C Minimale
Température           :  28°C Maximale
Biscarrosse
Index UV              :  7
Coefficient matin     :  81
Coefficient après-midi:  88
Pleine mer matin      :  05:57
Base mer matin        :  00:12
Pleine mer après-midi :  18:11
Basse mer après-midi  :  12:22
Hauteur de houle      :  0.9 m
Force du vent         :  5
Direction du vent     :  SW
```

- Le troisième exemple *sports* fournit les classements du rugby Top 14, foot Ligue 1 et golf :

```
Classement rugby top 14

 1 ASM Clermont                 5 point(s)
 2 Stade Français               5 point(s)
 3 Union Bordeaux-Bègles        5 point(s)
 4 Racing 92                    5 point(s)
 5 Stade Rochelais              4 point(s)
 6 Castres Olympique            4 point(s)
 7 Stade Toulousain             2 point(s)
 8 Lyon LOU                     2 point(s)
 9 Montpellier Hérault Rugby    1 point(s)
10 Grenoble                     0 point(s)
11 Rugby Club Toulonnais        0 point(s)
12 Section Paloise              0 point(s)
13 Perpignan                    0 point(s)
14 SU Agen                      0 point(s)

Classement foot ligue 1

 1 Paris-SG                     9 point(s)
 2 Dijon                        9 point(s)
 3 Lille                        7 point(s)
 4 Lyon                         6 point(s)
 5 Nîmes                        6 point(s)
 6 Reims                        6 point(s)
 7 Toulouse                     6 point(s)
 8 Saint-Etienne                5 point(s)
 9 Marseille                    4 point(s)
10 Monaco                       4 point(s)
11 Strasbourg                   4 point(s)
12 Montpellier                  4 point(s)
13 Rennes                       4 point(s)
14 Amiens                       3 point(s)
15 Bordeaux                     3 point(s)
16 Caen                         2 point(s)
17 Nantes                       1 point(s)
18 Nice                         1 point(s)
19 Angers                       0 point(s)
20 Guingamp                     0 point(s)

```