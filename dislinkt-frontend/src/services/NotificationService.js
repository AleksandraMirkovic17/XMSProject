import axios from 'axios';

const NOTIFICATION_API_BASE_URL = 'http://localhost:4200/notification';

class NotificationService{

    getNotificationsByUser(userId){
        return axios.get(NOTIFICATION_API_BASE_URL + '/user/' + userId);
    }
}

export default new NotificationService();