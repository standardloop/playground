db.createUser(
    {
        user: "root",
        pwd: "mypassword",
        roles: [
            {
                role: "readWrite",
                db: "playground"
            }
        ]
    }
);
