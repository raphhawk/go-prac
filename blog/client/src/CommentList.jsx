export const CommentList = ({ comments }) => {
  const renderedComments = Object.values(comments).map((comment) => {
    return (
      <li key={comment.id} className="ml-5">
        {". " + comment.content}
      </li>
    );
  });
  return <ul className="text-sm">{renderedComments}</ul>;
};
