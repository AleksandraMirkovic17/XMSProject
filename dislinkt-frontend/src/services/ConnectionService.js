import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:4200/';

class ConnectionService{
    GetConnectionDetail(idA, idB){
        console.log(USER_API_BASE_URL+"connection/"+ConvertObjectIdToHexId(idA)+"/detail/"+ConvertObjectIdToHexId(idB))
        return axios.get(USER_API_BASE_URL+"connection/"+ConvertObjectIdToHexId(idA)+"/detail/"+ConvertObjectIdToHexId(idB))
}
    Connect(idA, idB){
        return axios.post(USER_API_BASE_URL+"connection/friend",
            {"userIDa" : ConvertObjectIdToHexId(idA),
                "userIDb" : ConvertObjectIdToHexId(idB)})
    }

    DeleteFriend(idA, idB){
        return axios.post(USER_API_BASE_URL+"connection/remove-friend",{
            "userIDa": ConvertObjectIdToHexId(idA),
            "userIDb": ConvertObjectIdToHexId(idB)
        })
    }

    SendFriendRequest(idA, idB){
        return axios.post(USER_API_BASE_URL+"connection/friend-request",
            {"userIDa" : ConvertObjectIdToHexId(idA),
                "userIDb" : ConvertObjectIdToHexId(idB)})
    }

    UnsendFriendRequest(idA, idB){
        return axios.post(USER_API_BASE_URL + "connection/remove-friend-request",
            {"userIDa" : ConvertObjectIdToHexId(idA),
                "userIDb" : ConvertObjectIdToHexId(idB)})
    }

    DeclineFriendRequest(idA, idB){
        return axios.post(USER_API_BASE_URL + "connection/decline-request",
            {
                "userIDa" : ConvertObjectIdToHexId(idA),
                "userIDb" : ConvertObjectIdToHexId(idB)
            })
    }

    RemoveFriend(idA, idB){
        return axios.post(USER_API_BASE_URL + "connection/remove-friend", {
            "userIDa" : ConvertObjectIdToHexId(idA),
            "userIDb" : ConvertObjectIdToHexId(idB)
        })
    }

    GetFriends(id){
        return axios.get(USER_API_BASE_URL + "connection/user/"+ConvertObjectIdToHexId(id)+"/friends")
    }

    GetFriendRequests(id){
        console.log(USER_API_BASE_URL + "connection/user/" +ConvertObjectIdToHexId(id) + "/friend-requests")
        return axios.get(USER_API_BASE_URL + "connection/user/" +ConvertObjectIdToHexId(id) + "/friend-request")
    }

    GetRecommenadations(id){
        console.log(USER_API_BASE_URL+"connection/recommendation/" + ConvertObjectIdToHexId(id))
        return axios.get(USER_API_BASE_URL+"connection/recommendation/" + ConvertObjectIdToHexId(id))
    }

    GetBlockeds(id){
        return axios.get(USER_API_BASE_URL+"connection/user/" + ConvertObjectIdToHexId(id)+"/blockeds")
    }

    BlockUser(idA, idB){
        return axios.post(USER_API_BASE_URL+"connection/block", {
            "userIDa" : ConvertObjectIdToHexId(idA),
            "userIDb" : ConvertObjectIdToHexId(idB)
        })
    }

    UnblockUser(idA, idB){
        return axios.post(USER_API_BASE_URL+"connection/unblock", {
            "userIDa" : ConvertObjectIdToHexId(idA),
            "userIDb" : ConvertObjectIdToHexId(idB)
        })
    }
}

function ConvertObjectIdToHexId(objectId){
    if (objectId.includes("Object")) {
        var split = objectId.split(/["]/)[1]
    console.log(split)

    return split
    } else {
     return objectId
    }

}

export default new ConnectionService()