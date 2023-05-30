export const CommentList = ({ comments }) => {
  const renderedComments = Object.values(comments).map((comment) => {
    return (
      <li key={comment.id} className="ml-5 list-disc">
        {comment.status === "approved" ? (
          comment.content
        ) : (
          <span className="italic">{"comment " + comment.status}</span>
        )}
      </li>
    );
  });
  return <ul className="text-sm">{renderedComments}</ul>;
};
