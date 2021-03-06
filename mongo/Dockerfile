FROM mongo
# Copy files.
COPY mongo-seed.sh .
COPY init/. .
# Exec the bash script.
ENTRYPOINT ["sh", "/mongo-seed.sh"]