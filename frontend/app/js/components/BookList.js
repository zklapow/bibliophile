import BookItem from './BookItem';

export default class BookList extends React.Component {
    render() {
        var rows = [];

        for (var book of this.props.books) {
            rows.push(<BookItem book={book} />);
        }

        return (
            <ul>
                {rows}
            </ul>
        )
    }
}