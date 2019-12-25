defmodule Day03.Wire do
  alias Day03.{Parser, WireController}

  def run(puzzle) do
    {:ok, controller} = load_data_controller(puzzle)

    controller
    |> WireController.matches()
    |> Enum.map(fn {x, y, steps} ->
      {dist(x, y), {x, y, steps}}
    end)
    |> Enum.reduce(fn {s, point}, {s_acc, point_acc} ->
      if s < s_acc do
        {s, point}
      else
        {s_acc, point_acc}
      end
    end)
  end

  def optimize_signal(puzzle) do
    {:ok, controller} = load_data_controller(puzzle)

    controller
    |> WireController.matches()
    |> Enum.map(fn {x, y, steps} ->
      {steps, {x, y, steps}}
    end)
    |> Enum.reduce(fn {s, point}, {s_acc, point_acc} ->
      if s < s_acc do
        {s, point}
      else
        {s_acc, point_acc}
      end
    end)
  end

  defp load_data_controller(puzzle) do
    case WireController.start_link() do
      {:ok, controller} ->
        puzzle
        |> Parser.parse()
        |> Enum.map(fn wire ->
          Task.async(fn ->
            send_to(controller, wire)
          end)
        end)
        |> Enum.map(fn t ->
          Task.await(t)
        end)

        {:ok, controller}
      _ ->
        :error
    end
  end

  defp send_to(controller, wire) do
    wire
    |> Enum.reduce({0, 0, 0}, fn {direction, n}, {x, y, steps} ->
      case direction do
        :left ->
          positions(controller, :horizontal, &(x - &1), fn _ -> y end, {x - n, y}, n, steps)

        :right ->
          positions(controller, :horizontal, &(x + &1), fn _ -> y end, {x + n, y}, n, steps)

        :up ->
          positions(controller, :vertical, fn _ -> x end, &(y + &1), {x, y + n}, n, steps)

        :down ->
          positions(controller, :vertical, fn _ -> x end, &(y - &1), {x, y - n}, n, steps)
      end
    end)
  end

  defp positions(controller, dir, fn_x, fn_y, {final_x, final_y}, n, steps) do
    1..n
    |> Enum.each(fn i ->
      WireController.position(controller, dir, steps + i, fn_x.(i), fn_y.(i))
    end)

    {final_x, final_y, steps + n}
  end

  defp dist(x, y), do: abs(x) + abs(y)
end
