import { useDrop } from "react-dnd";
import { Task, ColumnType } from "../helpers/helpers";
import { DraggableTaskCard } from "./DraggableTaskCard";

type TodoColumnProps = {
  title: string;
  tasks: Task[];
  columnType: ColumnType;
  onToggleTask?: (taskId: string) => void;
  onDeleteTask?: (taskId: string) => void;
  onMoveTask?: (taskId: string, newStatus: ColumnType) => void;
  onEditTask?: (taskId: string) => void;
};

export function TodoColumn({
  title,
  tasks,
  columnType,
  onToggleTask,
  onDeleteTask,
  onMoveTask,
  onEditTask,
}: TodoColumnProps) {
  const [{ isOver, canDrop }, drop] = useDrop({
    accept: "task",
    drop: (item: { id: string; task: Task }) => {
      if (item.task.status !== columnType && onMoveTask) {
        onMoveTask(item.id, columnType);
      }
    },
    collect: (monitor) => ({
      isOver: monitor.isOver(),
      canDrop: monitor.canDrop(),
    }),
  });

  const getColumnColor = () => {
    switch (columnType) {
      case "todo":
        return "bg-blue-500";
      case "in_progress":
        return "bg-yellow-500";
      case "done":
        return "bg-green-500";
      default:
        return "bg-gray-500";
    }
  };

  return (
    <div
      ref={drop as any}
      className={`bg-gray-200 rounded-lg shadow-md p-4 w-72 mx-2 transition-all ${
        isOver && canDrop ? "bg-gray-300 ring-2 ring-blue-400" : ""
      }`}
    >
      {/* Header */}
      <div className={`h-12 rounded-t-lg ${getColumnColor()}`}>
        <h2 className="text-white text-lg font-semibold flex items-center justify-center h-full">
          {title} ({tasks.length})
        </h2>
      </div>

      {/* Task List */}
      <div className="mt-4 space-y-3">
        {tasks.length === 0 ? (
          <div className="text-center text-gray-500 py-8 border-2 border-dashed border-gray-300 rounded">
            {isOver && canDrop ? "Solte aqui" : "Nenhuma tarefa"}
          </div>
        ) : (
          tasks.map((task) => (
            <DraggableTaskCard
              key={task.id}
              task={task}
              onToggleTask={onToggleTask}
              onDeleteTask={onDeleteTask}
              onEditTask={onEditTask}
            />
          ))
        )}
      </div>
    </div>
  );
}
