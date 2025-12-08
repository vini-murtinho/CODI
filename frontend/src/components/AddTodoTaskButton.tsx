type AddTodoTaskButtonProps = {
  onClick: () => void;
};

export function AddTodoTaskButton({ onClick }: AddTodoTaskButtonProps) {
  return (
    <div
      className="bg-purple-600 p-2 rounded-lg cursor-pointer h-12 w-36 flex items-center justify-center text-white text-sm font-bold shadow-md hover:bg-purple-700 transition-colors mt-4"
      onClick={onClick}
    >
      <h2 className="">Adicionar tarefa</h2>
    </div>
  );
}
