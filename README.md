Différents exemples de web scraping en utilisant le framework Colly (Golang).

![Colly](/img/colly.jpg)

- Le premier exemple *freelance-info* récupère les tarifs journaliers des freelances :

![freelance](/img/freelance.png)

```json
{"Administrateur BD":540,"Administrateur ERP":510,"Administrateur produits":470,"Administrateur réseaux":390,"Administrateur système":430,"Analyste":450,"Analyste d'exploitation":340,"Analyste programmeur":390,"Analyste réalisateur":410,"Architecte":620,"Architecte réseaux":590,"Assistant à maîtrise d'ouvrage":570,"Auditeur":660,"Chef de projet":570,"Concepteur BD":430,"Concepteur multimédia":380,"Concepteur télématique":450,"Consultant":600,"Consultant fonctionnel":650,"Consultant technique":540,"Consultant technique et formateur":640,"Directeur de projet":760,"Directeur informatique":790,"Développeur":410,"Expert":650,"Formateur":510,"Infographiste":330,"Ingénieur d'exploitation":430,"Ingénieur d'études":440,"Ingénieur de production":470,"Ingénieur réseaux":480,"Ingénieur système":500,"Maquettiste PAO":320,"Pupitreur/Pilote":300,"Responsable d'exploitation":550,"Responsable maintenance":390,"Rédacteur technique":340,"Support utilisateurs":310,"Technicien d'exploitation":280,"Technicien micro / réseaux":270,"Webmaster":320}
```
- Le second exemple *meteo* fournit des prévisions et informations de différents sites météos et surf :

![Forecast with Colly](/img/meteo.png)

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

![Sport scraping with Colly](/img/balls.png)

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

Classement mondiale de golf

 1 Etats-Unis                Dustin Johnson            10.09 point(s)
 2 Etats-Unis                Brooks Koepka             10.00 point(s)
 3 Etats-Unis                Justin Thomas             9.60 point(s)
 4 Angleterre                Justin Rose               8.41 point(s)
 5 Espagne                   Jon Rahm                  7.59 point(s)
 6 Italie                    Francesco Molinari        7.27 point(s)
 7 Irlande du nord           Rory McIlroy              6.89 point(s)
 8 Etats-Unis                Rickie Fowler             6.62 point(s)
 9 Etats-Unis                Jordan Spieth             6.58 point(s)
10 Australie                 Jason Day                 6.38 point(s)
11 Angleterre                Tommy Fleetwood           6.00 point(s)
12 Etats-Unis                Bryson DeChambeau         5.53 point(s)
13 Etats-Unis                Patrick Reed              5.27 point(s)
14 Etats-Unis                Bubba Watson              5.11 point(s)
15 Suède                     Alexander Noren           5.05 point(s)
16 Angleterre                Paul Casey                4.79 point(s)
17 Etats-Unis                Webb Simpson              4.62 point(s)
18 Etats-Unis                Tony Finau                4.46 point(s)
19 Japon                     Hideki Matsuyama          4.39 point(s)
20 Etats-Unis                Xander Schauffele         4.35 point(s)
21 Suède                     Henrik Stenson            4.31 point(s)
22 Australie                 Marc Leishman             4.29 point(s)
23 Etats-Unis                Patrick Cantlay           4.18 point(s)
24 Etats-Unis                Phil Mickelson            4.01 point(s)
25 Angleterre                Tyrrell Hatton            4.00 point(s)
26 Etats-Unis                Tiger Woods               3.91 point(s)
27 Espagne                   Sergio Garcia             3.82 point(s)
28 Etats-Unis                Kevin Kisner              3.76 point(s)
29 Etats-Unis                Kyle Stanley              3.71 point(s)
30 Espagne                   Rafael Cabrera-Bello      3.61 point(s)
31 Etats-Unis                Matt Kuchar               3.48 point(s)
32 Angleterre                Ian Poulter               3.41 point(s)
33 Thaïlande                 Kiradech Aphibarnrat      3.21 point(s)
34 Etats-Unis                Brian Harman              3.21 point(s)
35 Afrique du Sud            Louis Oosthuizen          3.16 point(s)
36 Etats-Unis                Charley Hoffman           3.08 point(s)
37 Afrique du Sud            Branden Grace             2.90 point(s)
38 Australie                 Adam Scott                2.88 point(s)
39 Australie                 Cameron Smith             2.87 point(s)
40 Danemark                  Thorbjorn Olesen          2.81 point(s)
41 Etats-Unis                Kevin Na                  2.76 point(s)
42 Etats-Unis                Gary Woodland             2.75 point(s)
43 Japon                     Satoshi Kodaira           2.74 point(s)
44 Etats-Unis                Daniel Berger             2.64 point(s)
45 Angleterre                Matthew Fitzpatrick       2.63 point(s)
46 Etats-Unis                Pat Perez                 2.62 point(s)
47 Corée du Sud              Byeong-Hun An             2.58 point(s)
48 Etats-Unis                Zach Johnson              2.48 point(s)
49 Etats-Unis                Brandt Snedeker           2.41 point(s)
50 Etats-Unis                Luke List                 2.39 point(s)
```