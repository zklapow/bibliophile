import dispatcher from '../dispatcher/AppDispatcher';
import ActionTypes from '../constants/ActionTypes';

export default {
    reloadBooks() {
        $.ajax({
            "url": "/bibliophile/v1/books",
            "method": "get"
        }).done((data) => {
            console.log(data);
            dispatcher.dispatch({
                actionType: ActionTypes.RELOAD,
                data: data
            })
        })
    }
}