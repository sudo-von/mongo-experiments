ls -1 *.json | sed 's/.json$//' | while read col; do 
    mongoimport --host mongo --db experiments --authenticationDatabase admin --username VoN --password NoV --collection $col --type json --drop --file /$col.json --jsonArray; 
    echo "Imported " $col " collection successfully";
done