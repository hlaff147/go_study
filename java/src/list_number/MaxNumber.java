package list_number;

public class MaxNumber {
    public static void main(String[] args) {
        int[] numbers = {1, 2, 3, 4, 9, 5, 6};

        if (numbers.length == 0) {
            System.err.println("Array is empty");
            return;
        }

        int max = numbers[0];

        for (int num : numbers) {
            if (num > max) {
                max = num;
            }
        }

        System.out.printf("The maximum number is: %d\n", max);
    }
}
