FROM ubuntu:latest

RUN apt-get update && apt-get install -y bash 

# && apt-get install jq -y && apt-get install -y curl

COPY . .

RUN chmod +x ./script.sh

CMD ["bash", "script.sh"]