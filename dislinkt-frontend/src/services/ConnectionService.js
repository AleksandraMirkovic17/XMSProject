import axios from 'axios';

const USER_API_BASE_URL = 'http://localhost:4200/';

class ConnectionService{
    GetConnectionDetail(idA, idB){
        return axios.get(USER_API_BASE_URL+"connection/"+ConvertObjectIdToHexId(idA)+"/detail/"+ConvertObjectIdToHexId(idB))
}
    Connect(idA, idB){
        return axios.post(USER_API_BASE_URL+"connection/friend",
            {"userIDa" : ConvertObjectIdToHexId(idA),
                "userIDb" : ConvertObjectIdToHexId(idB)})
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

    GetFriends(id){
        return axios.get(USER_API_BASE_URL + "connection/user/"+ConvertObjectIdToHexId(id)+"/friends")
    }

    GetFriendRequests(id){
        console.log(USER_API_BASE_URL + "connection/user/" +ConvertObjectIdToHexId(id) + "/friend-requests")
        return axios.get(USER_API_BASE_URL + "connection/user/" +ConvertObjectIdToHexId(id) + "/friend-request")
    }
}

function ConvertObjectIdToHexId(objectId){
    var split = objectId.split(/["]/)[1]
    return split

}

export default new ConnectionService()