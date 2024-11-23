package types;

public class JavaTypes {
    public static void main(String[] args) {
        byte byteVar = 127;

        short shortVar = 32000;

        int intVar = 2147483647;

        long longVar = 9223372036854775807L;

        float floatVar = 3.14f;

        double doubleVar = 3.1415926535;

        char charVar = 'A';

        boolean booleanVar = true;

        String stringVar = "Hello, World!";

        int[] intArray = {1, 2, 3, 4, 5};

        Object objectVar = new Object();

        System.out.println("byteVar: " + byteVar);
        System.out.println("shortVar: " + shortVar);
        System.out.println("intVar: " + intVar);
        System.out.println("longVar: " + longVar);
        System.out.println("floatVar: " + floatVar);
        System.out.println("doubleVar: " + doubleVar);
        System.out.println("charVar: " + charVar);
        System.out.println("booleanVar: " + booleanVar);
        System.out.println("stringVar: " + stringVar);
        System.out.println("intArray: " + java.util.Arrays.toString(intArray));
        System.out.println("objectVar: " + objectVar.toString());
    }
}
