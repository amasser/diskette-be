[![Go Report Card](https://goreportcard.com/badge/github.com/getdiskette/diskette)](https://goreportcard.com/report/github.com/getdiskette/diskette) [![Build Status](https://travis-ci.org/getdiskette/diskette.svg?branch=master)](https://travis-ci.org/getdiskette/diskette)

# Diskette

[![Join the chat at https://gitter.im/getdiskette/diskette](https://badges.gitter.im/getdiskette/diskette.svg)](https://gitter.im/getdiskette/diskette?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

MongoDB + REST API + Authentication + Authorization + Mail Notifications

## Status

Under heavy development.

## Roadmap

- [x] REST

    - [x] GET
    ```bash
    # examples:
    http localhost:5025/collection/user
    http localhost:5025/collection/user?q='{"name":"Joe Doe"}'
    http localhost:5025/collection/user?q='{"name":{"$ne":"Joe Doe"}}'
    ```

    - [x] POST
    ```bash
    # example:
    http POST localhost:5025/collection/user name="Joe Doe" email=joe.doe@gmail.com
    ```

    - [x] PUT
    ```bash
    # example:
    http PUT localhost:5025/collection/user?q='{"name":"Joe Doe"}' \$set:='{"email":"jdoe@gmail.com"}'
    ```

    - [x] DELETE
    ```bash
    # example
    http DELETE localhost:5025/collection/user?q='{"name":"Joe Doe"}'
    ```

- [x] Authentication

    - [x] sign up
    ```bash
    # example
    http POST localhost:5025/user/signup \
        email=joe.doe@gmail.com password=abc \
        profile:='{"name": "Joe Doe", "language": "en" }'
    ```

    - [x] confirm sign up
    ```bash
    http POST localhost:5025/user/confirm token=<confirmation_token>
    ```

    - [x] sign in
    ```bash
    # example
    http POST localhost:5025/user/signin email=joe.doe@gmail.com password=abc
    ```

    - [x] forgot password
    ```bash
    # example
    http POST localhost:5025/user/forgot-password email=joe.doe@gmail.com
    ```

    - [x] reset password
    ```bash
    # example
    http POST localhost:5025/user/reset-password token=<reset_token> password=123
    ```

    - [x] sign out
    ```bash
    http POST localhost:5025/session/signout \
        X-Diskette-Session-Token:<session_token>
    ```

    - [x] change password
    ```bash
    http POST localhost:5025/session/change-password \
        X-Diskette-Session-Token:<session_token> \
        oldPassword=<old_password> newPassword=<new_password>
    ```

    - [x] change email
    ```bash
    http POST localhost:5025/session/change-email \
        X-Diskette-Session-Token:<session_token> \
        password=<password> newEmail=<newEmail>
    ```

    - [x] set profile
    ```bash
    http POST localhost:5025/session/set-profile \
        X-Diskette-Session-Token:<session_token> \
        profile:='{"name": "Joe Doe"}'
    ```

- [x] User Management

    - [x] get users
    ```bash
    http localhost:5025/admin/get-users?q=<query> X-Diskette-Session-Token:<session_token>
    ```

    - [x] create user
    ```bash
    http POST localhost:5025/admin/create-user X-Diskette-Session-Token:<session_token> \
        email="joe.doe@gmail.com" password="123" \
        roles:='["admin"]' profile:='{"name": "Joe Doe", "lang": "en"}'
    ```

    - [x] change user password
    ```bash
    http POST localhost:5025/admin/change-user-password X-Diskette-Session-Token:<session_token> \
        userId=<user_id> newPassword=<new_password>
    ```
    - [x] change user email
    ```bash
    http POST localhost:5025/admin/change-user-email X-Diskette-Session-Token:<session_token> \
        userId=<user_id> newEmail=<new_email>
    ```

    - [x] set user roles
    ```bash
    # example
    http POST localhost:5025/admin/set-user-roles X-Diskette-Session-Token:<session_token> \
        userId=<user_id> newRoles:='["customer"]'
    ```

    - [x] set user profile
    ```bash
    # example
    http POST localhost:5025/admin/set-user-profile X-Diskette-Session-Token:<session_token> \
        userId=<user_id> newProfile:='{"name": "Joe Doe Jr."}'
    ```

    - [x] remove users
    ```bash
    # example
    http DELETE localhost:5025/admin/remove-users X-Diskette-Session-Token:<session_token> \
        userIds:='["56bf19d65a1d18b704000001", "56be731d5a1d18accd000001"]'
    ```

    - [x] signout users
    ```bash
    # example
    http POST localhost:5025/admin/signout-users X-Diskette-Session-Token:<session_token> \
        userIds:='["56bf19d65a1d18b704000001", "56be731d5a1d18accd000001"]'
    ```

    - [x] suspend users
    ```bash
    # example
    http POST localhost:5025/admin/suspend-users X-Diskette-Session-Token:<session_token> \
        userIds:='["56bf19d65a1d18b704000001", "56be731d5a1d18accd000001"]'
    ```

    - [x] unsuspend users
    ```bash
    # example
    http POST localhost:5025/admin/unsuspend-users X-Diskette-Session-Token:<session_token> \
        userIds:='["56bf19d65a1d18b704000001", "56be731d5a1d18accd000001"]'
    ```

    - [x] remove unconfirmed users
    ```bash
    http DELETE localhost:5025/admin/remove-unconfirmed-users X-Diskette-Session-Token:<session_token>
    ```

    - [x] remove expired reset keys
    ```bash
    http POST localhost:5025/admin/remove-expired-reset-keys X-Diskette-Session-Token:<session_token>
    ```

- [ ] Authorization
    - [ ] Document level access control
    ```json
    // example
    {
        "blog-post": {
            "read": true,
            "create": "session.userId != null",
            "update": "session.userId === doc.authorId || 'admin' in session.userRoles",
            "remove": "session.userId === doc.authorId || 'admin' in session.userRoles"
        }
    }
    ```

- [ ] Mail Notifications:

    - [ ] send mail
    ```bash
    http POST localhost:5025/sendmail/<key>?q=<query>&l=<language_tag>
    ```

    - [ ] confirm signup
    ```bash
    # example
    http POST localhost:5025/sendmail/confirmSignup?q='{"email":"joe.doe@gmail.com"}'&l="en"
    ```

    - [ ] reset password
    ```bash
    # example
    http POST localhost:5025/sendmail/resetPassword?q='{"email":"joe.doe@gmail.com"}'&l="en"
    ```

- [ ] Javascript library for usage in the browser


## License

MIT
