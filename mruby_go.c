#include <mruby.h>
#include <mruby/compile.h>
#include <stdio.h>

static mrb_state *mrb;

static void mruby_open()
{
  mrb = mrb_open();
}

static void mruby_load_from_string(char *code)
{
  if (!mrb)
  {
    fprintf(stderr, "mrb_state is not open!");
    return;
  }
  mrb_load_string(mrb, code);
}

static void mruby_load_from_file(char *fname, char *entry_point_func_name)
{
  FILE *file = fopen(fname, "r");
  if (!mrb)
  {
    fprintf(stderr, "mrb_state is not open!");
    return;
  }
  mrb_value obj = mrb_load_file(mrb, file);
  mrb_funcall(mrb, obj, entry_point_func_name, 0);
}

static void mruby_close()
{
  if (!mrb)
  {
    fprintf(stderr, "mrb_state is not open!");
    return;
  }
  mrb_close(mrb);
}