const EDUCATIVE_LIVE_VM_URL = "ed-6307162657128448.educative.run"
const USER_API_BASE_URL = `http://${this.EDUCATIVE_LIVE_VM_URL}:8080`

class UserService {
    getUsers() {
        let users;
    
        fetch(`${USER_API_BASE_URL}/users`)
            .then((res) => {
                if (!res.ok) throw Error(res.status)
                return res.json()
            }).then((res) => {
                console.log({res})
                users = res.users
            }).catch((err) => console.log(err))
    
        return users
    }
    
    createUser(user) {
        let newUser;
    
        fetch(`${USER_API_BASE_URL}/users`, {
            body: { user: newUser }
        })
            .then((res) => {
                if (!res.ok) throw Error(res.status)
                return res.json()
            }).then((res) => {
                console.log({res})
                newUser = res.user
            }).catch((err) => console.log(err))
    }
    
    getUserById(userId) {
        let user;
    
        fetch(`${USER_API_BASE_URL}/users/${userId}`)
            .then((res) => {
                if (!res.ok) throw Error(res.status)
                return res.json()
            }).then((res) => {
                console.log({res})
                user = res.user
            }).catch((err) => console.log(err))
    }

    updateUser(user, userId) {

    }

    deleteUser(userId) {

    }
}

export default UserService
