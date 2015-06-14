import BookStore from '../stores/BookStore';
import BooksAction from '../actions/BooksAction';
import BookList from './BookList';

export default class App extends React.Component {

	constructor() {
	    super();
	    this.state = this.getState();
	}

	_onChange() {
		this.setState(this.getState());
	}

	getState() {
        console.log(BookStore);
		return {
			books: BookStore.get()
		};
	}

	componentDidMount() {
        this.subscription = BookStore.addOnChange(this._onChange.bind(this));
        BooksAction.reloadBooks();
	}

	componentWillUnmount() {
        this.subscription.remove();
	}

	render() {
		return (
            <div className="row">
                <div className="large-8 large-centered columns">
                    <BookList books={this.state.books}/>
                </div>
            </div>
	    )
	}
}
