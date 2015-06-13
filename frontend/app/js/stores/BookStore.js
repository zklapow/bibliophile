import dispatcher from '../dispatcher/AppDispatcher';
import ActionTypes from '../constants/ActionTypes';

var books = []
export default GeneralStore.define()
    .defineGet(() => { return books;})
    .defineResponseTo(ActionTypes.RELOAD, (data) => {
        books = data;
    })
    .register(dispatcher);
