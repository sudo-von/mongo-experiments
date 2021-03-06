FROM mongo
COPY mongo-seed.sh .
COPY mongo/init/. .
ENTRYPOINT ["sh", "/mongo-seed.sh"]