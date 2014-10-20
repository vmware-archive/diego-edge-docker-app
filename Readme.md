#Diego Edge Docker App

This is a docker app used in the Diego Edge Whetstone tests.

##Running Locally

    make
    docker run diego-edge
    
    
###To hit the app:
    boot2docker ssh #When using boot2docker
    curl localhost:8080


##To push to dockerhub
  
  
   docker push dajulia3/diego-edge-docker-app
