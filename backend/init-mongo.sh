mongosh -- "$MONGO_INITDB_DATABASE" <<EOF
    const rootUser = '$MONGO_INITDB_ROOT_USERNAME';
    const rootPassword = '$MONGO_INITDB_ROOT_PASSWORD';

    const adminDb = db.getSiblingDB('admin');
    adminDb.auth(rootUser, rootPassword);

    const user = '$MONGO_INITDB_USERNAME';
    const passwd = '$MONGO_INITDB_PASSWORD';

    const targetDbStr = '$MONGO_INITDB_DATABASE';
    const targetDb = db.getSiblingDB(targetDbStr);

    targetDb.createUser({
      user: user,
      pwd: passwd,
      roles: [
        {
          role: "readWrite",
          db: targetDbStr
        }
      ]
    });
EOF