import axios from 'axios';

const MESSAGE_API_BASE_URL = 'http://localhost:4200/message';

class MessageService{

    getMessagesByUser(userId){
        return axios.get(MESSAGE_API_BASE_URL + '/user/' + userId);
    }

    sendMessage(message){
        return axios.post(MESSAGE_API_BASE_URL, message);
    }
}

export default new MessageService();