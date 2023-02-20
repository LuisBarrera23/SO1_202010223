FROM ubuntu:latest

# COnfiguracion para la zona horaria
RUN apt-get update && apt-get install -y tzdata bash
#Fijamos la zona horaria 
RUN ln -sf /usr/share/zoneinfo/America/Guatemala/etc/localtime
#Se crea un entorno con la hora que debe de ser para Guatemala
RUN echo "export CURRENT_TIME=$(date '+%Y-%m-%d %H:%M:%S')" >> /root/.bashrc

# && apt-get install jq -y && apt-get install -y curl

COPY . .

RUN chmod +x ./script.sh

CMD ["/bin/bash", "-c", "source /root/.bashrc && ./script.sh"]