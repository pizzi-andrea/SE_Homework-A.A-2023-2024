echo "L'utente user_01 tenta di leggere il contenuto del file shadow"

cat /etc/shadow

echo "user_01 non avendo i permessi di root non può accedere al file delle password criptate, deve trovare una scorciatoglia ..."

ls .

echo " Il file main sarà la scorciatoglia... è possibile sfruttare un overflow del programma più il fatto di avere il setuid abilitato"

stat main

echo " E' possibile usare l'attacco progettato per aprire una shell con privileggi di root 

python3 inject.py
bash load.sh 




 