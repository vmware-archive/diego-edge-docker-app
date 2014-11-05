#Diego Edge Docker App

This is a docker app used in the Diego Edge Whetstone tests.

##Running Locally

    make
    docker run -p 8080:8080 diegoedge/test-app
    
    
###To hit the app:
    boot2docker ssh #When using boot2docker
    curl localhost:8080


##To push to dockerhub
    docker push diegoedge/test-app
