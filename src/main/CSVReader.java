
import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;

public class CSVReader {

  public static void main(String[] args) {

    String line = "";
    String cvsSplitBy = ",";
    File folder = new File("./../../input");

    try {

      while (true) {

        if (folder.listFiles().length > 0) {
          for (final File fileEntry : folder.listFiles()) {

            try (BufferedReader br = new BufferedReader(new FileReader(fileEntry))) {
              int count = 0;
              while ((line = br.readLine()) != null) {

                String[] dataLine = line.split(cvsSplitBy);
                if (dataLine != null && dataLine.length == 3) {

                  System.out.println(++count + " Country: '" + dataLine[0]
                      + "', \t\t Total Death: '" + dataLine[1] + "', \t\tDeath Per 100.000 person "
                      + dataLine[2]);
                }
              }
              br.close();
              fileEntry.delete();


            } catch (IOException e) {
              e.printStackTrace();
            }
          }
        }

      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }

}