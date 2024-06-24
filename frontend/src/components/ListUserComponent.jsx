import React, { useEffect } from 'react'
import { useState } from 'react'
import UserService from '../services/UserService'

const service = new UserService()

const ListUserComponent = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    const usersQuery = service.getUsers()
    setUsers(usersQuery)
  }, [])

  return <>
    <h2>Users List</h2>
    <button onClick={addUser}>Add User</button>
    <UserTable {...{users}} />
  </>
}

const UserTable = ({ users }) => {
  return (
    <table>
        <caption>
            Front-end web developer course 2021
        </caption>
        <thead>
            <tr>
                <th scope="col">Id</th>
                <th scope="col">First Name</th>
                <th scope="col">Middle Name</th>
                <th scope="col">Last Name</th>
                <th scope="col">Email</th>
                <th scope="col">Gender</th>
                <th scope="col">Civil Status</th>
                <th scope="col">Birthday</th>
                <th scope="col">Contact</th>
                <th scope="col">Address</th>
                <th scope="col">Age</th>
            </tr>
        </thead>
        <tbody>
            {users.map((u) => (
                <tr>
                    <td>{u.id}</td>
                    <td>{u.firstName}</td>
                    <td>{u.middleName}</td>
                    <td>{u.lastName}</td>
                    <td>{u.email}</td>
                    <td>{u.gender}</td>
                    <td>{u.civilStatus}</td>
                    <td>{u.birthday}</td>
                    <td>{u.contact}</td>
                    <td>{u.address}</td>
                    <td>{u.age}</td>
                    <button>Update</button>
                    <button>Delete</button>
                    <button>View</button>
                </tr>
            ))}
            
        </tbody>
    </table>
  )
}

function addUser() {
    window.location.href = "/users/new"
}

function editUser() {

}

function deleteUser() {
    if (window.confirm("Are you sure you want to delete this user?")) {
        
    }
}

export default ListUserComponent