import AppDispatcher from '../dispatcher/AppDispatcher';

export default class BaseStore extends EventEmitter {

  constructor() {
    super();
  }

  register(actionSubscribe) {
    this._dispatchToken = AppDispatcher.register(actionSubscribe);
  }

  get dispatchToken() {
    return this._dispatchToken;
  }

  emitChange() {
    this.emit('RELOAD');
  }

  addChangeListener(cb) {
    this.on('RELOAD', cb)
  }

  removeChangeListener(cb) {
    this.removeListener('RELOAD', cb);
  }
}
