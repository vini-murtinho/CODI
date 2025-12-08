import { useDrag } from "react-dnd";
import { Task } from "../helpers/helpers";

interface DraggableTaskCardProps {
  task: Task;
  onToggleTask?: (taskId: string) => void;
  onDeleteTask?: (taskId: string) => void;
  onEditTask?: (taskId: string) => void;
}

export function DraggableTaskCard({
  task,
  onToggleTask,
  onDeleteTask,
  onEditTask,
}: DraggableTaskCardProps) {
  const [{ isDragging }, drag] = useDrag({
    type: "task",
    item: { id: task.id, task },
    collect: (monitor) => ({
      isDragging: monitor.isDragging(),
    }),
  });

  const getBorderColor = () => {
    switch (task.status) {
      case "todo":
        return "border-blue-500";
      case "in_progress":
        return "border-yellow-500";
      case "done":
        return "border-green-500";
      default:
        return "border-gray-500";
    }
  };

  return (
    <div
      ref={drag as any}
      className={`bg-white p-3 rounded shadow border-l-4 cursor-move transition-all ${getBorderColor()} ${isDragging ? "opacity-50 transform rotate-2" : "hover:shadow-md"
        }`}
    >
      <div className="flex justify-between items-start">
        <div className="flex-1">
          <h3
            className={`font-medium ${task.completed ? "line-through text-gray-500" : "text-gray-800"
              }`}
          >
            {task.title}
          </h3>
          {task.description && (
            <p
              className={`text-sm mt-1 ${task.completed ? "text-gray-400" : "text-gray-600"
                }`}
            >
              {task.description}
            </p>
          )}
          <div className="mt-2">
            <span
              className={`inline-block px-2 py-1 text-xs rounded-full font-medium ${task.status === "todo"
                  ? "bg-blue-100 text-blue-600"
                  : task.status === "in_progress"
                    ? "bg-yellow-100 text-yellow-600"
                    : "bg-green-100 text-green-600"
                }`}
            >
              {task.status === "todo"
                ? "A Fazer"
                : task.status === "in_progress"
                  ? "Em Progresso"
                  : "Concluída"}
            </span>
          </div>
        </div>

        {/* Action buttons */}
        <div className="flex gap-1 ml-2">
          {/* Toggle completion */}
          {onToggleTask && (
            <button
              onClick={() => onToggleTask(task.id)}
              className={`p-1 rounded text-xs hover:bg-gray-100 transition-colors ${task.completed ? "bg-green-100 text-green-600" : "bg-gray-100 text-gray-600"
                }`}
              title={task.completed ? "Marcar como não concluída" : "Marcar como concluída"}
            >
              {task.completed ? "✅" : "⬜"}
            </button>
          )}

          {!task.completed && onEditTask && (
            <button
              onClick={() => onEditTask(task.id)}
              className="p-1 rounded text-xs bg-blue-100 text-blue-600 hover:bg-blue-200 transition-colors"
              title="Editar tarefa"
            >
              ✏️
            </button>
          )}

          {onDeleteTask && (
            <button
              onClick={() => onDeleteTask(task.id)}
              className="p-1 rounded text-xs bg-red-100 text-red-600 hover:bg-red-200 transition-colors"
              title="Deletar tarefa"
            >
              🗑️
            </button>
          )}
        </div>
      </div>
    </div>
  );
}
