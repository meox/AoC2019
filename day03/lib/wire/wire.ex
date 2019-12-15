defmodule Day03.Wire do
  alias Day03.{Parser, WireController}

  def run(puzzle) do
    {:ok, controller} = WireController.start_link()

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

    controller
    |> WireController.matches()
    |> Enum.map(fn {x, y} ->
      {dist(x, y), {x, y}}
    end)
    |> Enum.reduce(fn {s, point}, {s_acc, point_acc} ->
      if s < s_acc do
        {s, point}
      else
        {s_acc, point_acc}
      end
    end)
  end

  def send_to(controller, wire) do
    wire
    |> Enum.reduce({0, 0}, fn {direction, n}, {x, y} ->
      case direction do
        :left ->
          positions(controller, :horizontal, &(x - &1), fn _ -> y end, {x - n, y}, n)

        :right ->
          positions(controller, :horizontal, &(x + &1), fn _ -> y end, {x + n, y}, n)

        :up ->
          positions(controller, :vertical, fn _ -> x end, &(y + &1), {x, y + n}, n)

        :down ->
          positions(controller, :vertical, fn _ -> x end, &(y - &1), {x, y - n}, n)
      end
    end)
  end

  defp positions(controller, dir, fn_x, fn_y, final, n) do
    1..n
    |> Enum.each(fn i ->
      WireController.position(controller, dir, fn_x.(i), fn_y.(i))
    end)

    final
  end

  defp dist(x, y), do: abs(x) + abs(y)
end
