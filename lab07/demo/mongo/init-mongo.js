// use testDB;
// db.getCollection('test').find()

db = new Mongo().getDB("testDB");

db.createCollection('test', { capped: false });

db.test.insert([
    { "item": 1 },
    { "item": 2 },
    { "item": 3 },
    { "item": 4 },
    { "item": 5 }
]);
