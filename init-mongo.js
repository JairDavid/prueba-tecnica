// init-mongo.js
db = db.getSiblingDB('task-db'); // Selecciona la base de datos

db.createUser(
    {
        user: "admin",
        pwd: "D6Rhr2ey7aMzuTK75gcS",
        roles: [ { role: "readWrite", db: "task-db"} ],
        passwordDigestor: "server",
    }
)

db.createCollection('task'); // Crea la colección
