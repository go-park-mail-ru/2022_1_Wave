db.createUser(
        {
            user: "wave",
            pwd: "music",
            roles: [
                {
                    role: "readWrite",
                    db: "urls"
                }
            ]
        }
);
