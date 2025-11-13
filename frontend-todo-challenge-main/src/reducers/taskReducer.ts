import { Task, ColumnType } from "../helpers/helpers";

export type TaskAction =
  | { type: "ADD_TASK"; payload: Task }
  | { type: "TOGGLE_TASK"; payload: string }
  | { type: "DELETE_TASK"; payload: string }
  | { type: "UPDATE_TASK"; payload: Task }
  | { type: "SET_TASKS"; payload: Task[] }
  | { type: "MOVE_TASK"; payload: { taskId: string; newStatus: ColumnType } };

export interface TaskState {
  tasks: Task[];
}

export function taskReducer(state: TaskState, action: TaskAction): TaskState {
  switch (action.type) {
    case "ADD_TASK":
      return {
        ...state,
        tasks: [...state.tasks, action.payload],
      };

    case "TOGGLE_TASK":
      return {
        ...state,
        tasks: state.tasks.map((task) =>
          task.id === action.payload
            ? { ...task, completed: !task.completed }
            : task
        ),
      };

    case "DELETE_TASK":
      return {
        ...state,
        tasks: state.tasks.filter((task) => task.id !== action.payload),
      };

    case "UPDATE_TASK":
      return {
        ...state,
        tasks: state.tasks.map((task) =>
          task.id === action.payload.id ? action.payload : task
        ),
      };

    case "MOVE_TASK":
      return {
        ...state,
        tasks: state.tasks.map((task) =>
          task.id === action.payload.taskId
            ? {
              ...task,
              status: action.payload.newStatus,
              completed: action.payload.newStatus === "done",
            }
            : task
        ),
      };

    case "SET_TASKS":
      return {
        ...state,
        tasks: action.payload,
      };

    default:
      return state;
  }
}

// Estado inicial
export const initialTaskState: TaskState = {
  tasks: [],
};
