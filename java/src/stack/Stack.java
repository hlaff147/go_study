package stack ;

import java.util.ArrayList;

public class Stack {
    private ArrayList<Integer> stack = new ArrayList<>();

    public void push(int value) {
        stack.add(value);
    }

    public Integer pop() {
        if (stack.isEmpty()) {
            return null;
        } else {
            return stack.remove(stack.size() - 1);
        }
    }

    public static void main(String[] args) {
        Stack stack = new Stack();
        stack.push(10);
        stack.push(20);
        Integer val;
        if ((val = stack.pop()) != null) {
            System.out.println(val); // 20
        }
        if ((val = stack.pop()) != null) {
            System.out.println(val); // 10
        }
        if (stack.pop() == null) {
            System.out.println("Stack isEmpty!");
        }
    }
}

