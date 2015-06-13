
export default class BookItem extends React.Component {
    render() {
        return (
            <li>Title: {this.props.book.Title}</li>
        )
    }
}