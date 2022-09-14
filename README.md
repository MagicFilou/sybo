# sybo

//comments
gorm is powerfull -> less maintenance of the queries if the data layer changes
several drivers available for other sql techs

could have used direct call to db and quesries, it could be handy on a small project or wiht significant advance sql requirements

//be consisten with the scorem should be the same not score and another place "hightscore"

//maybe an update user might be an easier way to use than smaller update endpoints

//would recommend to use PATCH rather than PUT. PUT shoudl be to replace the entire object and PATCH only to do partial ones

//friend a comma separated list works, however it has advantages and inconvenient. Some thogh into that
//pros: easy to replace all, quick and simple solution. No need to check for the current list to delete the non existing ones
//cons: does not scale, cannot check as a foreign key.
//in the long term i would recommend to do the friends as a separated table but it was so much easier to start witha comma separated list
//Same kind of applies for scores, but

//make consts on all the items which are reused -> names, separators etc...

//return for get friends if there is no friends found will give the following results
