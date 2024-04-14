#!/bin/bash


clear # Elimina output dei precedenti comandi

# start server 
# go build ../cmd/webapi && go run webapi

# Test API doLogin()
username=("debian" "kali" "opensuse" "ubuntu" "fedora" "mangeia" "knopix" "garuda")
echo	"[Login Test]"
for u in "${username[@]}"
do
	curl -X 'POST' \
		'http://localhost:3000/session' \
		-H 'accept: application/json' \
		-H 'Content-Type: application/json' \
		-d '{"name": "'"$u"'"}'
	echo ":$u"
done
echo "[end Login test]"

# Test API followingUser()
echo "[start following test]"
for i in {1..10}
do
	from=$(expr 1 + $RANDOM % "${#username[@]}")
	to=$(expr 1 + $RANDOM % "${#username[@]}")

	if [ $from != $to ]
	then
		curl -X 'PUT' \
			"http://localhost:3000/users/$to/followers/$from" \
			-H 'accept: application/json' \
			-H "Authorization: Bearer $from"
	fi
done

echo "[end following test]"

echo "[start post photo]"

for i in {1..5}
do
	from=$(expr 1 + $RANDOM % "${#username[@]}")

	curl -X 'POST' \
	"http://localhost:3000/users/$from/myPhotos/" \
	-H 'accept: application/json' \
	-H 'Content-Type: multipart/form-data' \
	-H "Authorization: Bearer $from " \
	-F 'descr="Una foto spazialeeeee!" '\
	-F 'img=@tiny.png'
done

echo "[end post photo]"
echo "[start comment photo]"
for i in {1..5}
do
    
    from=$(expr 1 + $RANDOM % "${#username[@]}")
    to=$(expr 1 + $RANDOM % "${#username[@]}")
    photo=$(expr 1 + $RANDOM % 4)
    curl -X 'POST' \
	"http://localhost:3000/users/$to/myPhotos/$photo/comments/" \
	-H 'accept: application/json' \
	-H "Authorization: Bearer $from " \
	-H 'Content-Type: application/json' \
	-d "{
	\"author\": {
		\"username\": \"${username[$from]}\"
	},
	\"text\": \"😀 i like you photo! 😀\"
	}"
    
    
done

echo "[end comment photo]"

echo "[start like photo]"
for i in {1..4}
do
    
    from=$(expr 1 + $RANDOM % "${#username[@]}")
    to=$(expr 1 + $RANDOM % "${#username[@]}")
    photo=$(expr 1 + $RANDOM % 4)
    curl -X 'PUT' \
	"http://localhost:3000/users/$to/myPhotos/$photo/likes/$from"\
	-H 'accept: application/json' \
	-H "Authorization: Bearer $from"
    
    
done
echo "[end like photo]"


echo "[start list users]"

curl -X 'GET' \
	"http://localhost:3000/users/?limit=&username=" \
	-H 'accept: application/json' \
	-H "Authorization: Bearer 1"

for i in {1..5}
do
    
    from=$(expr 1 + $RANDOM % "${#username[@]}")
    usr=$(expr 1 + $RANDOM % "${#username[@]}")
    limit=$(expr 1 + $RANDOM % 10)
	curl -X 'GET' \
	"http://localhost:3000/users/?limit=$limit&username=${username[$usr]}" \
	-H 'accept: application/json' \
	-H "Authorization: Bearer $from"
    
    
done
echo "[end list users]"

echo "[start show user profile]"

for i in {1..5}
do
    
   
    usr=$(expr 1 + $RANDOM % "${#username[@]}")
	curl -X 'GET' \
		"http://localhost:3000/users/$usr/" \
		-H 'accept: application/json' \
		-H "Authorization: Bearer $usr"
    
done

echo "[end show user profile]"

echo "[start download photo]"



curl -X 'GET' \
	"http://localhost:3000/images/2" \
	-H 'accept: image/*' \
	-H "Authorization: Bearer 2" > out2.png


echo "[end download photo]"


echo "[start test mystrem]"


for i in {1..3}
do
    
   
    usr=$(expr 1 + $RANDOM % "${#username[@]}")

	curl -X 'GET' \
		"http://localhost:3000/users/$usr/myStream/?limit=50&sortBy=dataUpdate&reverse=desc" \
		-H 'accept: application/json' \
		-H "Authorization: Bearer $usr"
    
done

echo "[end test mystream]"
