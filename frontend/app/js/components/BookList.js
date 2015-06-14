import BookItem from './BookItem';

export default class BookList extends React.Component {
    render() {
        var rows = [];

        for (var book of this.props.books) {
            rows.push(
                <li>
                    <div className="panel clearfix">
                        <BookItem book={book} />
                    </div>
                </li>
            );
        }

        return (
            <div>
                <ul className="booklist">
                    {rows}
                </ul>
            </div>
        )
    }
}