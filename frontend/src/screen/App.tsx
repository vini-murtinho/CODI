import { useReducer, useState, useEffect } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import "../App.css";
import { AddTodoTaskButton } from "../components/AddTodoTaskButton";
import { TodoColumn } from "../components/TodoColumn";
import {
  fetchAllTasks,
  createTask,
  updateTask,
  deleteTask,
  getPendingTasks,
  getCompletedTasks,
  getInProgressTasks,
  ColumnType,
} from "../helpers/helpers";
import { AddTodoTaskModal } from "../components/AddTodoTaskModal";
import { taskReducer, initialTaskState } from "../reducers/taskReducer";

function App() {
  const [state, dispatch] = useReducer(taskReducer, initialTaskState);
  const [openModal, setOpenModal] = useState<boolean>(false);
  const [editingTaskId, setEditingTaskId] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true);

  // Carrega as tarefas do backend na primeira renderização
  useEffect(() => {
    const loadTasks = async () => {
      try {
        const tasks = await fetchAllTasks();
        dispatch({ type: "SET_TASKS", payload: tasks });
      } catch (error) {
        console.error("Erro ao carregar tarefas:", error);
      } finally {
        setLoading(false);
      }
    };

    loadTasks();
  }, []);

  const pendingTasks = getPendingTasks(state.tasks);
  const inProgressTasks = getInProgressTasks(state.tasks);
  const completedTasks = getCompletedTasks(state.tasks);

  const handleToggleTask = async (taskId: string) => {
    const task = state.tasks.find((t) => t.id === taskId);
    if (!task) return;

    const newStatus = task.status === "done" ? "todo" : "done";
    const updated = await updateTask(taskId, { status: newStatus as ColumnType });

    if (updated) {
      dispatch({ type: "MOVE_TASK", payload: { taskId, newStatus: newStatus as ColumnType } });
    }
  };

  const handleDeleteTask = async (taskId: string) => {
    const success = await deleteTask(taskId);
    if (success) {
      dispatch({ type: "DELETE_TASK", payload: taskId });
    }
  };

  const handleMoveTask = async (taskId: string, newStatus: ColumnType) => {
    const updated = await updateTask(taskId, { status: newStatus });
    if (updated) {
      dispatch({ type: "MOVE_TASK", payload: { taskId, newStatus } });
    }
  };

  const handleEditTask = (taskId: string) => {
    setEditingTaskId(taskId);
    setOpenModal(true);
  };

  const getEditingTask = () => {
    return editingTaskId
      ? state.tasks.find((task) => task.id === editingTaskId)
      : null;
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex flex-row p-8 gap-4 justify-center">
        <div className="bg-gray-200 rounded-lg shadow-md p-4 w-72 mx-2 animate-pulse" />
        <div className="bg-gray-200 rounded-lg shadow-md p-4 w-72 mx-2 animate-pulse" />
        <div className="bg-gray-200 rounded-lg shadow-md p-4 w-72 mx-2 animate-pulse" />
      </div>
    );
  }

  return (
    <DndProvider backend={HTML5Backend}>
      <div>
        <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex flex-row p-8 gap-4 justify-center">
          <AddTodoTaskButton onClick={() => setOpenModal(true)} />
          <TodoColumn
            title="A Fazer"
            tasks={pendingTasks}
            columnType="todo"
            onToggleTask={handleToggleTask}
            onDeleteTask={handleDeleteTask}
            onMoveTask={handleMoveTask}
            onEditTask={handleEditTask}
          />
          <TodoColumn
            title="Em Progresso"
            tasks={inProgressTasks}
            columnType="in_progress"
            onToggleTask={handleToggleTask}
            onDeleteTask={handleDeleteTask}
            onMoveTask={handleMoveTask}
            onEditTask={handleEditTask}
          />
          <TodoColumn
            title="Concluídas"
            tasks={completedTasks}
            columnType="done"
            onToggleTask={handleToggleTask}
            onDeleteTask={handleDeleteTask}
            onMoveTask={handleMoveTask}
            onEditTask={handleEditTask}
          />
        </div>
        {openModal && (
          <AddTodoTaskModal
            onClose={() => {
              setOpenModal(false);
              setEditingTaskId(null);
            }}
            onAddTask={async (task) => {
              if (editingTaskId) {
                // Editando tarefa existente
                const updated = await updateTask(editingTaskId, {
                  title: task.title,
                  description: task.description,
                });

                if (updated) {
                  dispatch({
                    type: "UPDATE_TASK",
                    payload: {
                      ...updated,
                      ...task,
                      id: editingTaskId,
                    },
                  });
                }
              } else {
                // Adicionando nova tarefa
                const newTask = await createTask(task.title, task.description);

                if (newTask) {
                  dispatch({ type: "ADD_TASK", payload: newTask });
                }
              }
              setOpenModal(false);
              setEditingTaskId(null);
            }}
            editingTask={getEditingTask()}
          />
        )}
      </div>
    </DndProvider>
  );
}

export default App;
