# For each file that ends with .json, the while loop will drop and insert the same file in the mongo db instance.
# The collection name will be the same as the file name without .json extension.
ls -1 *.json | sed 's/.json$//' | while read col; do 
    mongoimport --host mongo --db experiments --authenticationDatabase admin --username VoN --password NoV --collection $col --type json --drop --file /$col.json --jsonArray; 
    echo "Imported " $col " collection successfully";
done