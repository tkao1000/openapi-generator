package org.openapitools.codegen.templating.mustache;

import com.samskivert.mustache.Mustache;
import com.samskivert.mustache.Template;
import org.openapitools.codegen.CodegenConfig;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.io.Writer;

public class DefaultValueLambda implements Mustache.Lambda {
  private CodegenConfig generator = null;

  public DefaultValueLambda generator(final CodegenConfig generator) {
    this.generator = generator;
    return this;
  }

  @Override
  public void execute(Template.Fragment fragment, Writer writer) throws IOException {
    String text = fragment.execute();
    if (generator != null && text.contains("Default: null")) {
      System.out.println("DEFAULTVALUE IS EMPTY: ");
    } else {
      System.out.println("DEFAULTVALUE IS NOT EMPTY: ");
      writer.write("\n" + text);
    }
  }
}
