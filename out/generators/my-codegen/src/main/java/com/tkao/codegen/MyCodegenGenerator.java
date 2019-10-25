package com.tkao.codegen;

import org.openapitools.codegen.*;
import io.swagger.models.properties.*;

import java.util.*;
import java.io.File;

public class MyCodegenGenerator extends DefaultCodegen implements CodegenConfig {

  // source folder where to write the files
  protected String sourceFolder = "src";
  protected String apiVersion = "1.0.0";

  /**
   * Configures the type of generator.
   * 
   * @return  the CodegenType for this generator
   * @see     org.openapitools.codegen.CodegenType
   */
  public CodegenType getTag() {
    return CodegenType.OTHER;
  }

  /**
   * Configures a friendly name for the generator.  This will be used by the generator
   * to select the library with the -g flag.
   * 
   * @return the friendly name for the generator
   */
  public String getName() {
    return "my-codegen";
  }

  /**
   * Provides an opportunity to inspect and modify operation data before the code is generated.
   */
  @SuppressWarnings("unchecked")
  @Override
  public Map<String, Object> postProcessOperationsWithModels(Map<String, Object> objs, List<Object> allModels) {

    // to try debugging your code generator:
    // set a break point on the next line.
    // then debug the JUnit test called LaunchGeneratorInDebugger

    Map<String, Object> results = super.postProcessOperationsWithModels(objs, allModels);

    Map<String, Object> ops = (Map<String, Object>)results.get("operations");
    ArrayList<CodegenOperation> opList = (ArrayList<CodegenOperation>)ops.get("operation");

    // iterate over the operation and perhaps modify something
    for(CodegenOperation co : opList){
      // example:
      // co.httpMethod = co.httpMethod.toLowerCase();
    }

    return results;
  }

  /**
   * Returns human-friendly help for the generator.  Provide the consumer with help
   * tips, parameters here
   * 
   * @return A string value for the help message
   */
  public String getHelp() {
    return "Generates a my-codegen client library.";
  }

  public MyCodegenGenerator() {
    super();

    // set the output folder here
    outputFolder = "generated-code/my-codegen";

    /**
     * Models.  You can write model files using the modelTemplateFiles map.
     * if you want to create one template for file, you can do so here.
     * for multiple files for model, just put another entry in the `modelTemplateFiles` with
     * a different extension
     */
    modelTemplateFiles.put(
      "model.mustache", // the template to use
      ".sample");       // the extension for each file to write

    /**
     * Api classes.  You can write classes for each Api file with the apiTemplateFiles map.
     * as with models, add multiple entries with different extensions for multiple files per
     * class
     */
    apiTemplateFiles.put(
      "api.mustache",   // the template to use
      ".sample");       // the extension for each file to write

    /**
     * Template Location.  This is the location which templates will be read from.  The generator
     * will use the resource stream to attempt to read the templates.
     */
    templateDir = "my-codegen";

    /**
     * Api Package.  Optional, if needed, this can be used in templates
     */
    apiPackage = "org.openapitools.api";

    /**
     * Model Package.  Optional, if needed, this can be used in templates
     */
    modelPackage = "org.openapitools.model";

    /**
     * Reserved words.  Override this with reserved words specific to your language
     */
    reservedWords = new HashSet<String> (
      Arrays.asList(
        "sample1",  // replace with static values
        "sample2")
    );

    /**
     * Additional Properties.  These values can be passed to the templates and
     * are available in models, apis, and supporting files
     */
    additionalProperties.put("apiVersion", apiVersion);

    /**
     * Supporting Files.  You can write single files for the generator with the
     * entire object tree available.  If the input file has a suffix of `.mustache
     * it will be processed by the template engine.  Otherwise, it will be copied
     */
    supportingFiles.add(new SupportingFile("tfResourceTest.mustache",   // the input template or file
        "",                                                       // the destination folder, relative `outputFolder`
        "resourceTest.go")                                          // the output file
    );

    supportingFiles.add(new SupportingFile("tfResource.mustache",   // the input template or file
        "",                                                       // the destination folder, relative `outputFolder`
        "resource.go")                                          // the output file
    );

    supportingFiles.add(new SupportingFile("tfClient.mustache",   // the input template or file
        "",                                                       // the destination folder, relative `outputFolder`
        "client.go")                                          // the output file
    );

    supportingFiles.add(new SupportingFile("tfResourceImportTest.mustache",   // the input template or file
        "",                                                       // the destination folder, relative `outputFolder`
        "importResourceTest.go")                                          // the output file
    );


    /**
     * Language Specific Primitives.  These types will not trigger imports by
     * the client generator
     */
    /**
     * Language Specific Primitives.  These types will not trigger imports by
     * the client generator
     */
    languageSpecificPrimitives = new HashSet<String>(
        Arrays.asList(
            "string",
            "bool",
            "uint",
            "uint32",
            "uint64",
            "int",
            "int32",
            "int64",
            "float32",
            "float64",
            "complex64",
            "complex128",
            "rune",
            "byte"
        )
    );

    // type mappings for golang
    typeMapping = new HashMap<String, String>() {{
      put("integer", "int32");
      put("long", "int64");
      put("number", "float32");
      put("float", "float32");
      put("double", "float64");
      put("boolean", "bool");
      put("string", "string");
      put("UUID", "string");
      put("date", "string");
      put("DateTime", "time.Time");
      put("password", "string");
      put("File", "*os.File");
      put("file", "*os.File");
      // map binary to string as a workaround
      // the correct solution is to use []byte
      put("binary", "string");
      put("ByteArray", "string");
      put("object", "interface{}");
      put("UUID", "string");
      put("array", "[]");
      put("map", "map[string]");
    }};
  }

  /**
   * Escapes a reserved word as defined in the `reservedWords` array. Handle escaping
   * those terms here.  This logic is only called if a variable matches the reserved words
   * 
   * @return the escaped term
   */
  @Override
  public String escapeReservedWord(String name) {
    return "_" + name;  // add an underscore to the name
  }

  /**
   * Location to write model files.  You can use the modelPackage() as defined when the class is
   * instantiated
   */
  public String modelFileFolder() {
    return outputFolder + "/" + sourceFolder + "/" + modelPackage().replace('.', File.separatorChar);
  }

  /**
   * Location to write api files.  You can use the apiPackage() as defined when the class is
   * instantiated
   */
  @Override
  public String apiFileFolder() {
    return outputFolder + "/" + sourceFolder + "/" + apiPackage().replace('.', File.separatorChar);
  }

  /**
   * override with any special text escaping logic to handle unsafe
   * characters so as to avoid code injection
   *
   * @param input String to be cleaned up
   * @return string with unsafe characters removed or escaped
   */
  @Override
  public String escapeUnsafeCharacters(String input) {
    //TODO: check that this logic is safe to escape unsafe characters to avoid code injection
    return input;
  }

  /**
   * Escape single and/or double quote to avoid code injection
   *
   * @param input String to be cleaned up
   * @return string with quotation mark removed or escaped
   */
  public String escapeQuotationMark(String input) {
    //TODO: check that this logic is safe to escape quotation mark to avoid code injection
    return input.replace("\"", "\\\"");
  }
}