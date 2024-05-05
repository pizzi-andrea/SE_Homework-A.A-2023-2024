# SE_Homework

## Traccia 2 - SQL Injection

Realizzare un attacco di SQL injection di tipo inband (ovvero stesso canale utilizzato per l’injection della query malevola e per ricevere i risultati), basato su input dell’utente (ovvero l’attaccante inietta i comandi SQL fornendo un input opportunamente costruito) e che utilizzi una o più delle seguenti modalità: 

* Tautologia 
* commento di fine riga
*  query piggybacked. 

Mediante l’injection di opportuni comandi mostrare che e’ sia possibile compromettere almeno due delle proprietà CIA.

* Riservatezza
* Consistenza
* Disponibilità

### Suggerimenti:

Supponete che il sistema da attaccare sia un server su cui e’ installato un DBMS, su cui risiede il database e su cui risiede l’applicazione web vulnerabile (ad esempio una pagina Php che genera query al DB)
Il server può’ essere emulato mediante una macchina virtuale o un docker container
Scegliere versioni del software (OS, DB, Php, etc…) che siano vulnerabili ad attacchi SQLi
